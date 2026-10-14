package bson

import "bufio"

type Code struct {
	field
	value string
}

func readCode(r *bufio.Reader) (string, *Code, error) {
	f := Code{}

	n, err := f.read(r)
	if err != nil {
		return "", nil, err
	}

	return n, &f, nil
}

func (f *Code) Value() string { return f.value }
func (f *Code) Size() int     { return len(f.value) + 1 }

func (f *Code) read(r *bufio.Reader) (string, error) {
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
