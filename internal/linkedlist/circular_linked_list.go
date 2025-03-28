package linkedlist

import "fmt"

type CircularLinkedList[T comparable] struct {
	head *node[T]
	tail *node[T]
}

func (c *CircularLinkedList[T]) Add(value T) {
	newNode := &node[T]{
		value: value,
	}
	if c.tail == nil {
		c.head = newNode
		c.tail = newNode
		newNode.next = newNode
		newNode.prev = newNode
	} else {
		newNode.next = c.head
		newNode.prev = c.tail

		c.tail.next = newNode
		c.head.prev = newNode

		c.tail = newNode
	}
}

func (c *CircularLinkedList[T]) Show() string {
	var result string

	current := c.head
	for {
		result += fmt.Sprintf("%v -> ", current.value)
		current = current.next
		if current == c.head {
			break
		}
	}

	result += fmt.Sprintf("%v", c.head.value)
	return result
}

func (c *CircularLinkedList[T]) MultiAdd(values []T) {
	for _, v := range values {
		c.Add(v)
	}
}

func (c *CircularLinkedList[T]) IsExist(value T) bool {
	current := c.head
	for current.next != nil {
		if current.value == value {
			return true
		}
		current = current.next
		if current == c.head {
			break
		}
	}
	return false
}
