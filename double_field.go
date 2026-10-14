package bson

import (
	"bufio"
	"encoding/binary"
	"io"
	"math"
)

type DoubleField struct {
	field
	value float64
}

func ReadDoubleField(r *bufio.Reader) (string, *DoubleField, error) {
	f := DoubleField{}

	n, err := f.read(r)
	if err != nil {
		return "", nil, err
	}

	return n, &f, nil
}

func (f *DoubleField) Value() float64 {
	return f.value
}

func (f *DoubleField) read(r *bufio.Reader) (string, error) {
	n, err := f.readName(r)
	if err != nil {
		return "", err
	}

	v, err := f.readDouble(r)
	if err != nil {
		return "", err
	}

	f.value = v
	return n, nil
}

func (*DoubleField) readDouble(r io.Reader) (float64, error) {
	var data [8]byte
	if _, err := r.Read(data[:]); err != nil {
		return 0, err
	}

	return math.Float64frombits(binary.LittleEndian.Uint64(data[:])), nil
}
