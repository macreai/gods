package ds

import "github.com/macreai/gods/internal/linkedlist"

func NewSingleLinkedList[T comparable]() *linkedlist.SingleLinkedList[T] {
	return &linkedlist.SingleLinkedList[T]{}
}

func NewDoublyLinkedList[T comparable]() *linkedlist.DoublyLinkedList[T] {
	return &linkedlist.DoublyLinkedList[T]{}
}

func NewCircularLinkedList[T comparable]() *linkedlist.CircularLinkedList[T] {
	return &linkedlist.CircularLinkedList[T]{}
}
