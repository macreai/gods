package tree

import "golang.org/x/exp/constraints"

type node[T constraints.Ordered] struct {
	value T
	right *node[T]
	left  *node[T]
}
