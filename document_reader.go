package bson

import (
	"bufio"
	"errors"
	"io"
)

var (
	ErrUnknownElement = errors.New("unknown element")
)

func Read(r io.Reader) (*Document, error) {
	d := Document{}
	if err := d.read(bufio.NewReader(r)); err != nil {
		return nil, err
	}

	return &d, nil
}
