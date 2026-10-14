package data

import (
	"io"
)

func ReadBool[T ~bool](r io.Reader) (T, error) {
	var data [1]byte
	_, err := r.Read(data[:])
	return data[0] != 0, err
}

func WriteBool[T ~bool](v T, w io.Writer) error {
	var data [1]byte
	if v {
		data[0] = 1
	}
	_, err := w.Write(data[:])
	return err
}
