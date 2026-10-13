package bson

type StringField struct {
	field
	value string
}

func (f *StringField) Value() string {
	return f.value
}
