package bson

import (
	"bufio"
	"encoding/binary"
	"errors"
	"io"
)

var (
	ErrUnknownElement = errors.New("unknown element")
)

type DocumentReader struct {
	r *bufio.Reader
}

func NewDocumentReader(r io.Reader) *DocumentReader {
	return &DocumentReader{
		bufio.NewReader(r),
	}
}

func (dr *DocumentReader) Parse() (*Document, error) {
	size, err := dr.parseUint32()
	if err != nil {
		return nil, err
	}

	fields := make([]*Field[any], 0)

	for err != io.EOF {
		t, err := dr.parseType()
		if err != nil {
			return nil, err
		}

		f, err := dr.parseField(t)
		if err != nil {
			return nil, err
		}

		fields = append(fields, f)
	}

	return &Document{
		size:   size,
		fields: fields,
	}, nil
}

func (dr *DocumentReader) parseField(element Element) (*Field[any], error) {
	switch element {
	case ElementNone:
		return nil, io.EOF
	case ElementString:
		return dr.parseStringField()
	case ElementInt64:
		return dr.parseInt64Field()
	default:
		return nil, ErrUnknownElement
	}
}

func (dr *DocumentReader) parseStringField() (*Field[any], error) {
	n, err := dr.parseString()
	if err != nil {
		return nil, err
	}

	v, err := dr.parseString()
	if err != nil {
		return nil, err
	}

	return &Field[any]{
		name:  n,
		value: v,
	}, nil
}

func (dr *DocumentReader) parseInt64Field() (*Field[any], error) {
	n, err := dr.parseString()
	if err != nil {
		return nil, err
	}

	v, err := dr.parseInt64()
	if err != nil {
		return nil, err
	}

	return &Field[any]{
		name:  n,
		value: v,
	}, nil
}

func (dr *DocumentReader) parseString() (string, error) {
	return dr.r.ReadString(0)
}

func (dr *DocumentReader) parseType() (Element, error) {
	var data [1]byte
	_, err := dr.r.Read(data[:])
	return Element(data[0]), err
}

func (dr *DocumentReader) parseUint32() (uint32, error) {
	var data [4]byte
	_, err := dr.r.Read(data[:])
	return binary.LittleEndian.Uint32(data[:]), err
}

func (dr *DocumentReader) parseInt64() (int64, error) {
	var data [4]byte
	_, err := dr.r.Read(data[:])
	return int64(binary.LittleEndian.Uint64(data[:])), err
}
