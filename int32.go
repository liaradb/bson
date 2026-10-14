package bson

import (
	"bufio"
	"encoding/binary"
	"io"
)

type Int32 struct {
	field
	value int32
}

func readInt32(r *bufio.Reader) (string, *Int32, error) {
	f := Int32{}

	n, err := f.read(r)
	if err != nil {
		return "", nil, err
	}

	return n, &f, nil
}

func (f *Int32) Value() int32 { return f.value }
func (*Int32) Size() int      { return 4 }

func (f *Int32) read(r *bufio.Reader) (string, error) {
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

func (*Int32) readInt32(r io.Reader) (int32, error) {
	var data [4]byte
	_, err := r.Read(data[:])
	return int32(binary.LittleEndian.Uint32(data[:])), err
}
