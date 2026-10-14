package bson

import "bufio"

type Binary struct {
	field
	value []byte
}

func readBinary(r *bufio.Reader) (string, *Binary, error) {
	f := Binary{}

	n, err := f.read(r)
	if err != nil {
		return "", nil, err
	}

	return n, &f, nil
}

func (f *Binary) Value() []byte {
	return f.value
}

func (f *Binary) read(r *bufio.Reader) (string, error) {
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

	f.value = v[:len(v)-1]
	return n, nil
}
