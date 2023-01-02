package restpl

import (
	"encoding/json"
)

type Mutation struct {
	Fields     []string
	NullFields []string
}

func (m *Mutation) ReadJSON(data []byte) error {
	m.Fields = []string{}
	m.NullFields = []string{}

	var container map[string]interface{}
	err := json.Unmarshal(data, &container)
	if err != nil {
		return err
	}

	for key, value := range container {
		m.Fields = append(m.Fields, key)
		if value == nil {
			m.NullFields = append(m.NullFields, key)
		}
	}

	return nil
}

func (m *Mutation) HasField(f string) bool {
	for _, field := range m.Fields {
		if field == f {
			return true
		}
	}

	return false
}

func (m *Mutation) IsNullField(f string) bool {
	for _, field := range m.NullFields {
		if field == f {
			return true
		}
	}

	return false
}
