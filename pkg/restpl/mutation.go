package restpl

type Mutation struct {
	Fields     []string
	NullFields []string
}

func (m *Mutation) ReadJson() {
	m.Fields = []string{}
	m.NullFields = []string{}
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
