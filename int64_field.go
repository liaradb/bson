package bson

import (
	"bufio"
	"encoding/binary"
	"io"
)

type Int64Field struct {
	field
	value int64
}

func ReadInt64Field(r *bufio.Reader) (string, *Int64Field, error) {
	f := Int64Field{}

	n, err := f.read(r)
	if err != nil {
		return "", nil, err
	}

	return n, &f, nil
}

func (f *Int64Field) Value() int64 {
	return f.value
}

func (f *Int64Field) read(r *bufio.Reader) (string, error) {
	n, err := f.readName(r)
	if err != nil {
		return "", err
	}

	v, err := f.readInt64(r)
	if err != nil {
		return "", err
	}

	f.value = v
	return n, nil
}

func (*Int64Field) readInt64(r io.Reader) (int64, error) {
	var data [4]byte
	_, err := r.Read(data[:])
	return int64(binary.LittleEndian.Uint64(data[:])), err
}
