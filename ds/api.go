package ds

import (
	"github.com/macreai/gods/internal/linkedlist"
	"github.com/macreai/gods/internal/tree"
	"golang.org/x/exp/constraints"
)

func NewSingleLinkedList[T comparable]() *linkedlist.SingleLinkedList[T] {
	return &linkedlist.SingleLinkedList[T]{}
}

func NewDoublyLinkedList[T comparable]() *linkedlist.DoublyLinkedList[T] {
	return &linkedlist.DoublyLinkedList[T]{}
}

func NewCircularLinkedList[T comparable]() *linkedlist.CircularLinkedList[T] {
	return &linkedlist.CircularLinkedList[T]{}
}

func NewBinaryTree[T constraints.Ordered]() *tree.BinaryTree[T] {
	return &tree.BinaryTree[T]{}
}
