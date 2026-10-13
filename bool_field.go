package bson

import (
	"bufio"
	"io"
)

type BoolField struct {
	field
	value bool
}

func ReadBoolield(r *bufio.Reader) (string, *BoolField, error) {
	f := BoolField{}

	n, err := f.read(r)
	if err != nil {
		return "", nil, err
	}

	return n, &f, nil
}

func (f *BoolField) Value() bool {
	return f.value
}

func (f *BoolField) read(r *bufio.Reader) (string, error) {
	n, err := f.readName(r)
	if err != nil {
		return "", err
	}

	v, err := f.readBool(r)
	if err != nil {
		return "", err
	}

	f.value = v
	return n, nil
}

func (*BoolField) readBool(r io.Reader) (bool, error) {
	var data [1]byte
	_, err := r.Read(data[:])
	return data[0] != 0, err
}
