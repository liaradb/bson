package bson

import (
	"bufio"
	"encoding/binary"
	"io"
	"math"
)

type Double struct {
	field
	value float64
}

func readDouble(r *bufio.Reader) (string, *Double, error) {
	f := Double{}

	n, err := f.read(r)
	if err != nil {
		return "", nil, err
	}

	return n, &f, nil
}

func (f *Double) Value() float64 {
	return f.value
}

func (f *Double) read(r *bufio.Reader) (string, error) {
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

func (*Double) readDouble(r io.Reader) (float64, error) {
	var data [8]byte
	if _, err := r.Read(data[:]); err != nil {
		return 0, err
	}

	return math.Float64frombits(binary.LittleEndian.Uint64(data[:])), nil
}
