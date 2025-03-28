package linkedlist

import "fmt"

type SingleLinkedList[T comparable] struct {
	head *node[T]
}

func (s *SingleLinkedList[T]) Add(value T) {
	newNode := &node[T]{
		value: value,
	}

	if s.head == nil {
		s.head = newNode
	} else {
		current := s.head
		for current.next != nil {
			current = current.next
		}
		current.next = newNode
	}
}

func (s *SingleLinkedList[T]) Show() string {
	var result string
	current := s.head
	for current != nil {
		result += fmt.Sprintf("%v -> ", current.value)
		current = current.next
	}
	result += "nil"
	return result
}

func (s *SingleLinkedList[T]) MultiAdd(values []T) {
	for _, v := range values {
		s.Add(v)
	}
}

func (s *SingleLinkedList[T]) IsExist(value T) bool {
	current := s.head
	for current.next != nil {
		if current.value == value {
			return true
		}
		current = current.next
	}
	return false
}
