package restql

type MultiString struct {
	Condition map[string][]interface{}
	Key       string
}

type MultiInt struct {
	Condition map[string][]interface{}
	Key       string
}

type MultiTime struct {
	Condition map[string][]interface{}
	Key       string
}
