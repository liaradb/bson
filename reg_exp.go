package bson

import (
	"bufio"

	"github.com/liaradb/bson/internal/data"
)

type RegExp struct {
	field
	pattern string
	options string
}

func readRegExp(r *bufio.Reader) (string, *RegExp, error) {
	f := RegExp{}

	n, err := f.read(r)
	if err != nil {
		return "", nil, err
	}

	return n, &f, nil
}

func (f *RegExp) Value() (string, string) { return f.pattern, f.options }
func (f *RegExp) Size() int               { return len(f.pattern) + len(f.options) + 2 }

func (f *RegExp) read(r *bufio.Reader) (string, error) {
	n, err := f.readName(r)
	if err != nil {
		return "", err
	}

	if _, err := f.readSize(r); err != nil {
		return "", err
	}

	p, err := data.ReadString(r)
	if err != nil {
		return "", err
	}

	o, err := data.ReadString(r)
	if err != nil {
		return "", err
	}

	f.pattern = p
	f.options = o
	return n, nil
}
