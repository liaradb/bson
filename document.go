package bson

import "iter"

type Document struct {
	size   uint32
	fields map[string]Field
}

func (d *Document) Fields() iter.Seq2[string, Field] {
	return func(yield func(string, Field) bool) {
		for n, f := range d.fields {
			if !yield(n, f) {
				return
			}
		}
	}
}
