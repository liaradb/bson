package bson

import "bufio"

type Document struct {
	size   uint32
	fields map[string]Field
}

type Field interface {
	_field()
}

type field struct{}

func (field) _field()

func (field) readName(r *bufio.Reader) (string, error) {
	return r.ReadString(0)
}
