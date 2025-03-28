package linkedlist

import "fmt"

type DoublyLinkedList[T comparable] struct {
	head *node[T]
	tail *node[T]
}

func (d *DoublyLinkedList[T]) Add(value T) {
	newNode := &node[T]{
		value: value,
	}

	if d.tail == nil {
		d.head = newNode
		d.tail = newNode
	} else {
		d.tail.next = newNode
		newNode.prev = d.tail
		d.tail = newNode
	}
}

func (d *DoublyLinkedList[T]) Show() (string, string) {
	var forwardLinkedList string
	var backwardLinkedList string

	forwardLinkedList = "nil -> "
	backwardLinkedList = "nil -> "

	head := d.head
	tail := d.tail

	for head != nil {
		forwardLinkedList += fmt.Sprintf("%v -> ", head.value)
		backwardLinkedList += fmt.Sprintf("%v -> ", tail.value)
		head = head.next
		tail = tail.prev
	}

	forwardLinkedList += "nil"
	backwardLinkedList += "nil"

	return forwardLinkedList, backwardLinkedList
}

func (d *DoublyLinkedList[T]) MultiAdd(values []T) {
	for _, v := range values {
		d.Add(v)
	}
}

func (d *DoublyLinkedList[T]) IsExist(value T) bool {
	current := d.head
	for current.next != nil {
		if current.value == value {
			return true
		}
		current = current.next
	}
	return false
}
