package data

import (
	"encoding/binary"
	"io"
)

func ReadInt64[T ~int64 | ~uint64](r io.Reader) (T, error) {
	var data [8]byte
	_, err := r.Read(data[:])
	return T(binary.LittleEndian.Uint64(data[:])), err
}

func WriteInt64[T ~int64 | ~uint64](v T, w io.Writer) error {
	var data [8]byte
	binary.LittleEndian.PutUint64(data[:], uint64(v))
	_, err := w.Write(data[:])
	return err
}
