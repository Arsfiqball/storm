package restql

import (
	"fmt"
	"net/url"
	"reflect"
	"strconv"
	"time"
)

const TagName = "restql"

func Decode(b []byte, out interface{}) error {
	qs := string(b)
	values, err := url.ParseQuery(qs)
	if err != nil {
		return err
	}

	t := reflect.TypeOf(out).Elem()

	if t.Kind() != reflect.Struct {
		return fmt.Errorf("out type is not struct")
	}

	v := reflect.ValueOf(out).Elem()
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)
		tag := field.Tag.Get(TagName)

		if len(tag) > 0 {
			switch field.Type {
			case reflect.TypeOf(MultiString{}):
				decoded := DecodeMultiString(values, tag)
				value.Set(reflect.ValueOf(decoded))

			case reflect.TypeOf(MultiInt{}):
				decoded, err := DecodeMultiInt(values, tag)
				if err != nil {
					return err
				}

				value.Set(reflect.ValueOf(decoded))

			case reflect.TypeOf(MultiTime{}):
				decoded, err := DecodeMultiTime(values, tag)
				if err != nil {
					return err
				}

				value.Set(reflect.ValueOf(decoded))
			}
		}
	}

	return nil
}

var validSuffixes = []string{
	"eq", "ne", "gt", "gte", "lt", "lte",
	"contain", "ncontain", "contains", "ncontains",
	"in", "nin",
}

func DecodeMultiString(values map[string][]string, name string) MultiString {
	conds := map[string][]interface{}{}

	for _, suffix := range validSuffixes {
		conds[suffix] = []interface{}{}
		key := name + "_" + suffix

		if val, ok := values[key]; ok {
			for _, v := range val {
				conds[suffix] = append(conds[suffix], v)
			}
		}
	}

	return MultiString{
		Key:       name,
		Condition: conds,
	}
}

func DecodeMultiInt(values map[string][]string, name string) (MultiInt, error) {
	conds := map[string][]interface{}{}

	for _, suffix := range validSuffixes {
		conds[suffix] = []interface{}{}
		key := name + "_" + suffix

		if val, ok := values[key]; ok {
			for _, v := range val {
				vi, err := strconv.Atoi(v)
				if err != nil {
					return MultiInt{}, err
				}

				conds[suffix] = append(conds[suffix], vi)
			}
		}
	}

	return MultiInt{
		Key:       name,
		Condition: conds,
	}, nil
}

func DecodeMultiTime(values map[string][]string, name string) (MultiTime, error) {
	conds := map[string][]interface{}{}

	for _, suffix := range validSuffixes {
		conds[suffix] = []interface{}{}
		key := name + "_" + suffix

		if val, ok := values[key]; ok {
			for _, v := range val {
				vi, err := time.Parse(time.RFC3339, v)
				if err != nil {
					return MultiTime{}, err
				}

				conds[suffix] = append(conds[suffix], vi)
			}
		}
	}

	return MultiTime{
		Key:       name,
		Condition: conds,
	}, nil
}
