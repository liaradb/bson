package bson

import "bufio"

type String struct {
	field
	value string
}

func readString(r *bufio.Reader) (string, *String, error) {
	f := String{}

	n, err := f.read(r)
	if err != nil {
		return "", nil, err
	}

	return n, &f, nil
}

func (f *String) Value() string { return f.value }
func (f *String) Size() int     { return len(f.value) + 1 }

func (f *String) read(r *bufio.Reader) (string, error) {
	n, err := f.readName(r)
	if err != nil {
		return "", err
	}

	if _, err := f.readSize(r); err != nil {
		return "", err
	}

	v, err := f.readString(r)
	if err != nil {
		return "", err
	}

	f.value = v
	return n, nil
}

func (*String) readString(r *bufio.Reader) (string, error) {
	v, err := r.ReadBytes(0)
	if err != nil {
		return "", err
	}

	return string(v[:len(v)-1]), nil
}
