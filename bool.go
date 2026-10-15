package bson

import (
	"bufio"

	"github.com/liaradb/bson/internal/data"
)

type Bool struct {
	field
	value bool
}

func readBool(r *bufio.Reader) (string, *Bool, error) {
	f := Bool{}

	n, err := f.read(r)
	if err != nil {
		return "", nil, err
	}

	return n, &f, nil
}

func (f *Bool) Value() bool { return f.value }
func (*Bool) Size() int     { return 1 }

func (f *Bool) read(r *bufio.Reader) (string, error) {
	n, err := f.readName(r)
	if err != nil {
		return "", err
	}

	v, err := data.ReadBool[bool](r)
	if err != nil {
		return "", err
	}

	f.value = v
	return n, nil
}
