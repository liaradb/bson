package bson

import (
	"bufio"
	"io"
	"iter"
)

type Document struct {
	field
	size   uint32
	fields map[string]Field
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
