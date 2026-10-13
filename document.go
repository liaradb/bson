package bson

type Document struct {
	size   uint32
	fields map[string]Field
}

type Field interface {
	_field()
}

type field struct{}

func (field) _field()
