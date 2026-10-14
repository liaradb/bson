package bson

import (
	"bufio"
	"io"
)

type Element int8

const (
	ElementNone          Element = 0
	ElementDouble        Element = 1   // 64-bit binary floating point
	ElementString        Element = 2   // UTF-8 string
	ElementDocument      Element = 3   // Embedded document
	ElementArray         Element = 4   // Array. See below for more information.
	ElementBinary        Element = 5   // Binary data
	ElementUndefined     Element = 6   // (value). Deprecated.
	ElementObjectID      Element = 7   // ObjectId
	ElementBool          Element = 8   // Boolean
	ElementDatetime      Element = 9   // UTC datetime. int64 is UTC milliseconds since the Unix epoch.
	ElementNull          Element = 10  // value
	ElementRegExp        Element = 11  // Regular expression. See below for more information.
	ElementDBPointer     Element = 12  // (byte*12)	DBPointer. Deprecated.
	ElementCode          Element = 13  // JavaScript code
	ElementSymbol        Element = 14  // Symbol. Deprecated.
	ElementCodeWithScope Element = 15  // JavaScript code with scope. Deprecated.
	ElementInt32         Element = 16  // 32-bit integer
	ElementTimestamp     Element = 17  // Timestamp. A special internal type used by MongoDB replication and sharding. The first 4 bytes are an increment and the second 4 bytes are a timestamp.
	ElementInt64         Element = 18  // 64-bit integer
	ElementDecimal128    Element = 19  // 128-bit decimal floating point
	ElementMinKey        Element = -1  // A special type that compares lower than all other possible BSON element values.
	ElementMaxKey        Element = 127 // A special type that compares higher than all other possible BSON element values.
)

func readElement(r *bufio.Reader) (Element, error) {
	var data [1]byte
	_, err := r.Read(data[:])
	return Element(data[0]), err
}

func (e Element) readField(r *bufio.Reader) (string, Field, error) {
	switch e {
	case ElementDouble:
		return readDouble(r)
	case ElementString:
		return readString(r)
	case ElementDocument:
		return readDocument(r)
	case ElementArray:
		return readArray(r)
	case ElementBinary:
		return readBinary(r)
	case ElementUndefined:
		return readNull(r)
	case ElementObjectID:
		return "", nil, ErrUnknownElement
	case ElementBool:
		return readBool(r)
	case ElementDatetime:
		return readDatetime(r)
	case ElementNull:
		return readNull(r)
	case ElementRegExp:
		return readRegExp(r)
	case ElementDBPointer:
		return "", nil, ErrUnknownElement
	case ElementCode:
		return readCode(r)
	case ElementSymbol:
		return readSymbol(r)
	case ElementCodeWithScope:
		return readCodeWithScope(r)
	case ElementInt32:
		return readInt32(r)
	case ElementTimestamp:
		return readTimestamp(r)
	case ElementInt64:
		return readInt64(r)
	case ElementDecimal128:
		return "", nil, ErrUnknownElement
	case ElementMinKey:
		return "", nil, ErrUnknownElement
	case ElementMaxKey:
		return "", nil, ErrUnknownElement
	case ElementNone:
		return "", nil, io.EOF
	default:
		return "", nil, ErrUnknownElement
	}
}
