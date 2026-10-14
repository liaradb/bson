package bson

import (
	"bufio"
	"io"
	"iter"
)

type Array struct {
	field
	size     uint32
	elements []Field
}

func readArray(r *bufio.Reader) (string, *Array, error) {
	a := Array{}
	n, err := a.readName(r)
	if err != nil {
		return "", nil, err
	}

	if err := a.read(r); err != nil {
		return "", nil, err
	}

	return n, &a, nil
}

func (a *Array) read(r *bufio.Reader) error {
	size, err := a.readSize(r)
	if err != nil {
		return err
	}

	elements := make([]Field, 0)

	for err != io.EOF {
		_, f, err := readField(r)
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		elements = append(elements, f)
	}

	a.size = size
	a.elements = elements
	return nil
}

func (a *Array) Fields() iter.Seq2[int, Field] {
	return func(yield func(int, Field) bool) {
		for n, f := range a.elements {
			if !yield(n, f) {
				return
			}
		}
	}
}
