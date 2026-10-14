package data

import "bufio"

func ReadString(r *bufio.Reader) (string, error) {
	v, err := r.ReadBytes(0)
	if err != nil {
		return "", err
	}

	return string(v[:len(v)-1]), nil
}

func WriteString(v string, w *bufio.Writer) error {
	_, err := w.WriteString(v)
	return err
}
