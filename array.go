package bson

import (
	"bufio"
	"io"
	"iter"
	"strconv"
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

func (a *Array) Size() int {
	size := 0
	for _, f := range a.elements {
		size += f.Size()
	}
	return size
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

func (a *Array) Get(n int) (Field, bool) {
	if n >= len(a.elements) {
		return nil, false
	}

	return a.elements[n], true
}

func (a *Array) getField(names []string) (Field, bool) {
	if len(names) == 0 {
		return nil, false
	}

	n, err := strconv.Atoi(names[0])
	if err != nil {
		return nil, false
	}

	f := a.elements[n]

	switch f := f.(type) {
	case *Document:
		return f.getField(names[1:])
	case *Array:
		return f.getField(names[1:])
	default:
		return f, true
	}
}
