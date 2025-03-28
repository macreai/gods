package linkedlist

type node[T comparable] struct {
	value T
	next  *node[T]
	prev  *node[T]
}
