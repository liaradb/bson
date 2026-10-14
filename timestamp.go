package bson

import (
	"bufio"
	"encoding/binary"
	"io"
)

type Timestamp struct {
	field
	counter   uint32
	timestamp uint32
}

func readTimestamp(r *bufio.Reader) (string, *Timestamp, error) {
	f := Timestamp{}

	n, err := f.read(r)
	if err != nil {
		return "", nil, err
	}

	return n, &f, nil
}

func (f *Timestamp) Value() (uint32, uint32) { return f.counter, f.timestamp }
func (*Timestamp) Size() int                 { return 8 }

func (f *Timestamp) read(r *bufio.Reader) (string, error) {
	n, err := f.readName(r)
	if err != nil {
		return "", err
	}

	c, err := f.readUint32(r)
	if err != nil {
		return "", err
	}

	t, err := f.readUint32(r)
	if err != nil {
		return "", err
	}

	f.counter = c
	f.timestamp = t
	return n, nil
}

func (*Timestamp) readUint32(r io.Reader) (uint32, error) {
	var data [4]byte
	_, err := r.Read(data[:])
	return binary.LittleEndian.Uint32(data[:]), err
}
