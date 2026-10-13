package bson

import (
	"bytes"
	"testing"
)

func TestDocumentReader(t *testing.T) {
	t.Parallel()

	dr := NewDocumentReader(bytes.NewBuffer(merge(
		[]byte{0x16, 0x00, 0x00, 0x00}, // total document size
		[]byte{0x02},                   // 0x02 = type String
		stringSlice("hello"),           // field name
		[]byte{0x06, 0x00, 0x00, 0x00}, //
		stringSlice("world"),           // field value
		[]byte{0x00},                   // 0x00 = type EOO ('end of object')
	)))

	d, err := dr.Parse()
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
