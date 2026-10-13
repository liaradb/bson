package bson

import "bufio"

type StringField struct {
	field
	value string
}

func (f *StringField) Value() string {
	return f.value
}

func (f *StringField) read(r *bufio.Reader) (string, error) {
	n, err := f.readName(r)
	if err != nil {
		return "", err
	}

	if _, err := f.readSize(r); err != nil {
		return "", err
	}

	v, err := r.ReadBytes(0)
	if err != nil {
		return "", err
	}

	f.value = string(v[:len(v)-1])
	return n, nil
}
