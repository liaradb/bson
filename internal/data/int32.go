package data

import (
	"encoding/binary"
	"io"
)

func ReadInt32[T ~int32 | ~uint32](r io.Reader) (T, error) {
	var data [4]byte
	_, err := r.Read(data[:])
	return T(binary.LittleEndian.Uint32(data[:])), err
}

func WriteInt32[T ~int64 | ~uint64](v T, w io.Writer) error {
	var data [8]byte
	binary.LittleEndian.PutUint64(data[:], uint64(v))
	_, err := w.Write(data[:])
	return err
}
