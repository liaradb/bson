package bson

type Document struct {
	size   uint32
	fields []*Field[any]
}

type Field[T any] struct {
	name  string
	value T
}
