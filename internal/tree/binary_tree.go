package tree

import "golang.org/x/exp/constraints"

type BinaryTree[T constraints.Ordered] struct {
	root *node[T]
}

func (b *BinaryTree[T]) Add(value T) {
	newNode := &node[T]{
		value: value,
	}

	if b.root == nil {
		b.root = newNode
		return
	}

	current := b.root
	for {
		if value < current.value {
			if current.left == nil {
				current.left = newNode
				return
			}
			current = current.left
		} else {
			if current.right == nil {
				current.right = newNode
				return
			}
			current = current.right
		}

	}

}

func (b *BinaryTree[T]) MultiAddd(values []T) {
	for _, v := range values {
		b.Add(v)
	}
}

func (b *BinaryTree[T]) IsExist(value T) bool {
	if b.root == nil {
		return false
	}

	current := b.root

	for current != nil {
		if value == current.value {
			return true
		}

		if value < current.value {
			if current.left == nil {
				return false
			}
			current = current.left
		} else {
			if current.right == nil {
				return false
			}
			current = current.right
		}
	}

	return false
}
