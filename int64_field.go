package bson

type Int64Field struct {
	field
	value int64
}

func (f *Int64Field) Value() int64 {
	return f.value
}
