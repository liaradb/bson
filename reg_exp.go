package bson

import "bufio"

type RegExp struct {
	field
	pattern string
	options string
}

func readRegExp(r *bufio.Reader) (string, *RegExp, error) {
	f := RegExp{}

	n, err := f.read(r)
	if err != nil {
		return "", nil, err
	}

	return n, &f, nil
}

func (f *RegExp) Value() (string, string) {
	return f.pattern, f.options
}

func (f *RegExp) read(r *bufio.Reader) (string, error) {
	n, err := f.readName(r)
	if err != nil {
		return "", err
	}

	if _, err := f.readSize(r); err != nil {
		return "", err
	}

	p, err := f.readString(r)
	if err != nil {
		return "", err
	}

	o, err := f.readString(r)
	if err != nil {
		return "", err
	}

	f.pattern = p
	f.options = o
	return n, nil
}

func (*RegExp) readString(r *bufio.Reader) (string, error) {
	v, err := r.ReadBytes(0)
	if err != nil {
		return "", err
	}

	return string(v[:len(v)-1]), nil
}
