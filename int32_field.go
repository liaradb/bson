package bson

import (
	"bufio"
	"encoding/binary"
	"io"
)

type Int32Field struct {
	field
	value int32
}

func ReadInt32Field(r *bufio.Reader) (string, *Int32Field, error) {
	f := Int32Field{}

	n, err := f.read(r)
	if err != nil {
		return "", nil, err
	}

	return n, &f, nil
}

func (f *Int32Field) Value() int32 {
	return f.value
}

func (f *Int32Field) read(r *bufio.Reader) (string, error) {
	n, err := f.readName(r)
	if err != nil {
		return "", err
	}

	v, err := f.readInt32(r)
	if err != nil {
		return "", err
	}

	f.value = v
	return n, nil
}

func (*Int32Field) readInt32(r io.Reader) (int32, error) {
	var data [4]byte
	_, err := r.Read(data[:])
	return int32(binary.LittleEndian.Uint32(data[:])), err
}
