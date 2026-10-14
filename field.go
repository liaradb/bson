package bson

import (
	"bufio"
	"encoding/binary"
	"io"
)

type Field interface {
	_field()
}

type field struct{}

func (field) _field() {}

func (field) readName(r *bufio.Reader) (string, error) {
	v, err := r.ReadBytes(0)
	if err != nil {
		return "", err
	}

	return string(v[:len(v)-1]), nil
}

func (field) readSize(r io.Reader) (uint32, error) {
	var data [4]byte
	_, err := r.Read(data[:])
	return binary.LittleEndian.Uint32(data[:]), err
}

func readField(r *bufio.Reader) (string, Field, error) {
	e, err := readElement(r)
	if err != nil {
		return "", nil, err
	}

	return readElementField(e, r)
}

func readElement(r *bufio.Reader) (Element, error) {
	var data [1]byte
	_, err := r.Read(data[:])
	return Element(data[0]), err
}

func readElementField(element Element, r *bufio.Reader) (string, Field, error) {
	switch element {
	case ElementDouble:
		return ReadDoubleField(r)
	case ElementString:
		return ReadStringField(r)
	case ElementDocument:
		return ReadDocumentField(r)
	case ElementArray:
		return ReadArrayField(r)
	case ElementBool:
		return ReadBoolield(r)
	case ElementInt32:
		return ReadInt32Field(r)
	case ElementInt64:
		return ReadInt64Field(r)
	case ElementNone:
		return "", nil, io.EOF
	default:
		return "", nil, ErrUnknownElement
	}
}
