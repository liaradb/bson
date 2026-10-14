package data

import (
	"encoding/binary"
	"io"
	"math"
)

func ReadFloat64[T ~float64](r io.Reader) (T, error) {
	var data [8]byte
	if _, err := r.Read(data[:]); err != nil {
		return 0, err
	}

	return T(math.Float64frombits(binary.LittleEndian.Uint64(data[:]))), nil
}

func WriteFloat64[T ~float64](v T, w io.Writer) error {
	var data [8]byte
	binary.LittleEndian.PutUint64(data[:], math.Float64bits(float64(v)))
	_, err := w.Write(data[:])
	return err
}
