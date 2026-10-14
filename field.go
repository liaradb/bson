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

	return e.readField(r)
}
