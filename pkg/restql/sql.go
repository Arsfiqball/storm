package restql

import (
	"strconv"
	"strings"
)

type SqlExpression struct {
	Statement string
	Values    map[string]interface{}
}

func CombineSqlExpression(expressions ...SqlExpression) SqlExpression {
	var statements []string
	values := map[string]interface{}{}

	for _, expression := range expressions {
		if len(expression.Values) > 0 {
			statements = append(statements, expression.Statement)

			for name, val := range expression.Values {
				values[name] = val
			}
		}
	}

	return SqlExpression{
		Statement: strings.Join(statements, " AND "),
		Values:    values,
	}
}

func ConditionToSqlExpression(key string, conds map[string][]interface{}) SqlExpression {
	var statements []string
	values := map[string]interface{}{}
	count := 0

	addValue := func(op string, val interface{}) {
		count++
		argName := key + strconv.Itoa(count)
		argToken := "@" + argName
		statement := key + " " + op + " " + argToken
		statements = append(statements, statement)
		values[argName] = val
	}

	addValues := func(op string, vals []interface{}) {
		for _, v := range vals {
			addValue("=", v)
		}
	}

	basicCompMap := map[string]string{
		"eq":  "=",
		"ne":  "<>",
		"gt":  ">",
		"gte": ">=",
		"lt":  "<",
		"lte": "<=",
	}

	for w, s := range basicCompMap {
		if vals, ok := conds[w]; ok {
			addValues(s, vals)
		}
	}

	subStringCompMap := map[string]string{
		"contain":   "ILIKE",
		"ncontain":  "NOT ILIKE",
		"contains":  "LIKE",
		"ncontains": "NOT LIKE",
	}

	for w, s := range subStringCompMap {
		if val, ok := conds[w]; ok {
			for _, v := range val {
				if sv, ok := v.(string); ok {
					addValue(s, "%"+sv+"%")
				}
			}
		}
	}

	if val, ok := conds["in"]; ok && len(conds["in"]) > 0 {
		addValue("IN", val)
	}

	if val, ok := conds["nin"]; ok && len(conds["nin"]) > 0 {
		addValue("NOT IN", val)
	}

	return SqlExpression{
		Statement: strings.Join(statements, " AND "),
		Values:    values,
	}
}
