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

func ReadDocumentField(r *bufio.Reader) (string, *Document, error) {
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
		t, err := d.parseType(r)
		if err != nil {
			return err
		}

		n, f, err := d.parseField(t, r)
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

func (d *Document) parseType(r *bufio.Reader) (Element, error) {
	var data [1]byte
	_, err := r.Read(data[:])
	return Element(data[0]), err
}

func (d *Document) parseField(element Element, r *bufio.Reader) (string, Field, error) {
	switch element {
	case ElementDouble:
		return ReadDoubleField(r)
	case ElementString:
		return ReadStringField(r)
	case ElementDocument:
		return ReadDocumentField(r)
	case ElementArray:
		return ReadDocumentField(r)
	case ElementBool:
		return ReadBoolield(r)
	case ElementInt32:
		return ReadInt32Field(r)
	case ElementInt64:
		return ReadInt64Field(r)
	case ElementNone:
		return "", nil, io.EOF
	default:
		return "", nil, ErrUnknownElement
	}
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
