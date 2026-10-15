package bson

import (
	"bufio"

	"github.com/liaradb/bson/internal/data"
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

func (f *Double) Value() float64 { return f.value }
func (*Double) Size() int        { return 8 }

func (f *Double) read(r *bufio.Reader) (string, error) {
	n, err := f.readName(r)
	if err != nil {
		return "", err
	}

	v, err := data.ReadFloat64[float64](r)
	if err != nil {
		return "", err
	}

	f.value = v
	return n, nil
}
