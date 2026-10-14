package bson

import (
	"bytes"
	"fmt"
	"testing"
)

func TestDocumentReader(t *testing.T) {
	t.Parallel()

	d, err := Read(bytes.NewBuffer(merge(
		[]byte{0x16, 0x00, 0x00, 0x00}, // total document size
		[]byte{0x02},                   // 0x02 = type String
		stringSlice("hello"),           // field name
		[]byte{0x06, 0x00, 0x00, 0x00}, //
		stringSlice("world"),           // field value
		[]byte{0x00},                   // 0x00 = type EOO ('end of object')
	)))
	if err != nil {
		t.Fatal(err)
	}

	for n, f := range d.Fields() {
		if n != "hello" {
			t.Errorf("incorrect name: %v", n)
		}
		switch f := f.(type) {
		case *StringField:
			if v := f.Value(); v != "world" {
				t.Errorf("incorrect value: %v", v)
			}
		default:
			t.Fatal("incorrect type")
		}
		break
	}
}

func TestDocumentReader_Document(t *testing.T) {
	t.Parallel()

	d, err := Read(bytes.NewBuffer(merge(
		[]byte{0x31, 0x00, 0x00, 0x00},
		[]byte{0x04},
		stringSlice("BSON"),
		[]byte{0x26, 0x00, 0x00, 0x00},
		[]byte{0x02}, stringSlice("0"), []byte{0x08, 0x00, 0x00, 0x00}, stringSlice("awesome"),
		[]byte{0x01}, stringSlice("1"), []byte{0x33, 0x33, 0x33, 0x33, 0x33, 0x33, 0x14, 0x40},
		[]byte{0x10}, stringSlice("2"), []byte{0xc2, 0x07, 0x00, 0x00},
		[]byte{0x00},
		[]byte{0x00},
	)))
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(d)
}

func merge(b ...[]byte) []byte {
	d := make([]byte, 0)
	for _, i := range b {
		d = append(d, i...)
	}
	return d
}

func stringSlice(s string) []byte {
	return append([]byte(s), 0)
}
