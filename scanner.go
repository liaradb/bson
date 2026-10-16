package bson

import (
	"bufio"
	"errors"
	"io"
	"strings"

	"github.com/liaradb/bson/internal/data"
)

type Scanner struct {
	r       *bufio.Reader
	results map[string]any
}

func NewScanner(r io.Reader) *Scanner {
	return &Scanner{
		r:       bufio.NewReader(r),
		results: map[string]any{},
	}
}

func (s *Scanner) Read(name string) error {
	names := strings.Split(name, ".")
	e, err := data.ReadByte[Element](s.r)
	if err != nil {
		return err
	}

	switch e {
	case ElementDocument:
		s.readDocument(name, names)
	}

	return nil
}

func (s *Scanner) readDocument(name string, names []string) error {
	_, err := data.ReadInt32[uint32](s.r)
	if err != nil {
		return err
	}

	e, err := data.ReadByte[Element](s.r)
	if err != nil {
		return err
	}

	for e != ElementNone {
		n, err := data.ReadString(s.r)
		if err != nil {
			return err
		}

		if n == names[0] {
			return s.readElement(e, name, names[1:])
		}

		if err := s.skipElement(e); err != nil {
			return err
		}

		e, err = data.ReadByte[Element](s.r)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *Scanner) skipElement(e Element) error {
	size, err := s.getElementSize(e)
	if err != nil {
		return err
	}

	_, err = s.r.Discard(int(size))
	return err
}

func (s *Scanner) getElementSize(e Element) (int32, error) {
	switch e {
	case ElementDouble:
		return 8, nil
	case ElementString:
		return s.readSize()
	case ElementDocument:
		return s.readSize()
	case ElementArray:
		return s.readSize()
	case ElementBinary:
		return s.readSize()
	case ElementUndefined:
		return 0, nil
	case ElementObjectID:
		return 12, nil
	case ElementBool:
		return 1, nil
	case ElementDatetime:
		return 8, nil
	case ElementNull:
		return 0, nil
	case ElementRegExp:
		return s.readSize()
	case ElementDBPointer:
		return s.readSize()
	case ElementCode:
		return s.readSize()
	case ElementSymbol:
		return s.readSize()
	case ElementCodeWithScope:
		return s.readSize()
	case ElementInt32:
		return 4, nil
	case ElementTimestamp:
		return 8, nil
	case ElementInt64:
		return 8, nil
	case ElementDecimal128:
		return 16, nil
	case ElementMinKey:
		return 8, nil
	case ElementMaxKey:
		return 0, nil
	default:
		return 0, ErrUnknownElement
	}
}

func (s *Scanner) readSize() (int32, error) {
	return data.ReadInt32[int32](s.r)
}

var (
	ErrInvalidData = errors.New("invalid data")
)

func (s *Scanner) readElement(e Element, name string, names []string) error {
	switch e {
	case ElementDouble:
		return s.readDouble(names, name)
	case ElementString:
		return nil
	case ElementDocument:
		return nil
	case ElementArray:
		return nil
	case ElementBinary:
		return nil
	case ElementUndefined:
		return nil
	case ElementObjectID:
		return nil
	case ElementBool:
		return nil
	case ElementDatetime:
		return nil
	case ElementNull:
		return nil
	case ElementRegExp:
		return nil
	case ElementDBPointer:
		return nil
	case ElementCode:
		return nil
	case ElementSymbol:
		return nil
	case ElementCodeWithScope:
		return nil
	case ElementInt32:
		return nil
	case ElementTimestamp:
		return nil
	case ElementInt64:
		return nil
	case ElementDecimal128:
		return nil
	case ElementMinKey:
		return nil
	case ElementMaxKey:
		return nil
	default:
		return ErrUnknownElement
	}
}

func (s *Scanner) readDouble(names []string, name string) error {
	if len(names) != 0 {
		return ErrInvalidData
	}

	v, err := data.ReadFloat64[float64](s.r)
	if err != nil {
		return err
	}

	s.results[name] = v
	return nil
}
