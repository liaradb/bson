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

	fields := make(map[string]Field)

	for err != io.EOF {
		t, err := dr.parseType()
		if err != nil {
			return nil, err
		}

		n, f, err := dr.parseField(t)
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}

		fields[n] = f
	}

	return &Document{
		size:   size,
		fields: fields,
	}, nil
}

func (dr *DocumentReader) parseField(element Element) (string, Field, error) {
	switch element {
	case ElementString:
		return ReadStringField(dr.r)
	case ElementInt64:
		return ReadInt64Field(dr.r)
	case ElementInt32:
		return ReadInt32Field(dr.r)
	case ElementBool:
		return ReadBoolield(dr.r)
	case ElementNone:
		return "", nil, io.EOF
	default:
		return "", nil, ErrUnknownElement
	}
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
