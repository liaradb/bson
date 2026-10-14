package bson

import (
	"bufio"
	"encoding/binary"
	"io"
	"time"
)

type Datetime struct {
	field
	value time.Time
}

func readDatetime(r *bufio.Reader) (string, *Datetime, error) {
	f := Datetime{}

	n, err := f.read(r)
	if err != nil {
		return "", nil, err
	}

	return n, &f, nil
}

func (f *Datetime) Value() time.Time {
	return f.value
}

func (f *Datetime) read(r *bufio.Reader) (string, error) {
	n, err := f.readName(r)
	if err != nil {
		return "", err
	}

	v, err := f.readInt64(r)
	if err != nil {
		return "", err
	}

	f.value = v
	return n, nil
}

func (*Datetime) readInt64(r io.Reader) (time.Time, error) {
	var data [8]byte
	_, err := r.Read(data[:])
	return time.UnixMilli(int64(binary.LittleEndian.Uint64(data[:]))), err
}
