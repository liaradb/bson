package data

import "bufio"

func ReadBytes(r *bufio.Reader) ([]byte, error) {
	v, err := r.ReadBytes(0)
	if err != nil {
		return nil, err
	}

	return v[:len(v)-1], nil
}

func WriteBytes(v []byte, w *bufio.Writer) error {
	_, err := w.Write(v)
	return err
}
