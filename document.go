package bson

import (
	"bufio"
	"errors"
	"io"
	"iter"
	"strings"
)

var (
	ErrUnknownElement = errors.New("unknown element")
)

type Document struct {
	field
	size   uint32
	fields map[string]Field
}

func Read(r io.Reader) (*Document, error) {
	d := Document{}
	if err := d.read(bufio.NewReader(r)); err != nil {
		return nil, err
	}

	return &d, nil
}

func readDocument(r *bufio.Reader) (string, *Document, error) {
	d := Document{}
	n, err := d.readName(r)
	if err != nil {
		return "", nil, err
	}

	if err := d.read(r); err != nil {
		return "", nil, err
	}

	return n, &d, nil
}

func (d *Document) read(r *bufio.Reader) error {
	size, err := d.readSize(r)
	if err != nil {
		return err
	}

	fields := make(map[string]Field)

	for err != io.EOF {
		n, f, err := readField(r)
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		fields[n] = f
	}

	d.size = size
	d.fields = fields
	return nil
}

func (d *Document) Fields() iter.Seq2[string, Field] {
	return func(yield func(string, Field) bool) {
		for n, f := range d.fields {
			if !yield(n, f) {
				return
			}
		}
	}
}

func (d *Document) Get(n string) (Field, bool) {
	names := strings.Split(n, ".")
	return d.getField(names)
}

func (d *Document) getField(names []string) (Field, bool) {
	if len(names) == 0 {
		return nil, false
	}

	f, ok := d.fields[names[0]]
	if !ok {
		return nil, false
	}

	switch f := f.(type) {
	case *Document:
		return f.getField(names[1:])
	case *Array:
		return f.getField(names[1:])
	default:
		return f, true
	}
}
