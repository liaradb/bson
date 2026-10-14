package bson

import "bufio"

type CodeWithScope struct {
	field
	value string
}

func readCodeWithScope(r *bufio.Reader) (string, *CodeWithScope, error) {
	f := CodeWithScope{}

	n, err := f.read(r)
	if err != nil {
		return "", nil, err
	}

	return n, &f, nil
}

func (f *CodeWithScope) Value() string {
	return f.value
}

func (f *CodeWithScope) read(r *bufio.Reader) (string, error) {
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
