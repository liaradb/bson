package bson

import (
	"bufio"

	"github.com/liaradb/bson/internal/data"
)

type Int64 struct {
	field
	value int64
}

func readInt64(r *bufio.Reader) (string, *Int64, error) {
	f := Int64{}

	n, err := f.read(r)
	if err != nil {
		return "", nil, err
	}

	return n, &f, nil
}

func (f *Int64) Value() int64 { return f.value }
func (*Int64) Size() int      { return 8 }

func (f *Int64) read(r *bufio.Reader) (string, error) {
	n, err := f.readName(r)
	if err != nil {
		return "", err
	}

	v, err := data.ReadInt64[int64](r)
	if err != nil {
		return "", err
	}

	f.value = v
	return n, nil
}
