package bson

import "bufio"

type Null struct {
	field
}

func readNull(r *bufio.Reader) (string, *Null, error) {
	f := Null{}

	n, err := f.readName(r)
	if err != nil {
		return "", nil, err
	}

	return n, &f, nil
}
