package test

import (
	"testing"

	"github.com/macreai/gods/ds"
	"github.com/stretchr/testify/assert"
)

func TestSingleLL(t *testing.T) {
	numSLL := ds.NewSingleLinkedList[int]()
	strSLL := ds.NewSingleLinkedList[string]()

	assert.NotNil(t, numSLL)
	assert.NotNil(t, strSLL)
}

func TestAddSingleLL(t *testing.T) {
	numSLL := ds.NewSingleLinkedList[int]()
	strSLL := ds.NewSingleLinkedList[string]()

	numSLL.Add(19)
	strSLL.Add("Hello")
	assert.Equal(t, "19 -> nil", numSLL.Show())
	assert.Equal(t, "Hello -> nil", strSLL.Show())
}

func TestMultiAddSingleLL(t *testing.T) {
	numSLL := ds.NewSingleLinkedList[int]()
	strSLL := ds.NewSingleLinkedList[string]()

	numSLL.MultiAdd([]int{1, 9, 2, 3, 4, 5})
	strSLL.MultiAdd([]string{"Hello!", "I", "am", "just", "testing"})
	assert.Equal(t, "1 -> 9 -> 2 -> 3 -> 4 -> 5 -> nil", numSLL.Show())
	assert.Equal(t, "Hello! -> I -> am -> just -> testing -> nil", strSLL.Show())
}

func TestIsExistSingleLL(t *testing.T) {
	numSLL := ds.NewSingleLinkedList[int]()
	strSLL := ds.NewSingleLinkedList[string]()

	numSLL.MultiAdd([]int{1, 9, 2, 3, 4, 5})
	strSLL.MultiAdd([]string{"Hello!", "I", "am", "just", "testing"})

	assert.True(t, numSLL.IsExist(4))
	assert.True(t, strSLL.IsExist("am"))

	assert.False(t, numSLL.IsExist(19))
	assert.False(t, strSLL.IsExist("My"))
}

func TestDoublyLL(t *testing.T) {
	numDLL := ds.NewDoublyLinkedList[int]()
	strDLL := ds.NewDoublyLinkedList[string]()

	assert.NotNil(t, numDLL)
	assert.NotNil(t, strDLL)
}

func TestAddDoublyLL(t *testing.T) {
	numDLL := ds.NewDoublyLinkedList[int]()
	strDLL := ds.NewDoublyLinkedList[string]()

	numDLL.Add(19)
	strDLL.Add("Hello")

	forward, backward := numDLL.Show()
	assert.Equal(t, "nil -> 19 -> nil", forward)
	assert.Equal(t, "nil -> 19 -> nil", backward)

	forward, backward = strDLL.Show()
	assert.Equal(t, "nil -> Hello -> nil", forward)
	assert.Equal(t, "nil -> Hello -> nil", backward)
}

func TestMultiAddDoublyLL(t *testing.T) {
	numDLL := ds.NewDoublyLinkedList[int]()
	strDLL := ds.NewDoublyLinkedList[string]()

	numDLL.MultiAdd([]int{1, 9, 2, 3, 4, 5})
	strDLL.MultiAdd([]string{"Hello!", "I", "am", "just", "testing"})

	forward, backward := numDLL.Show()
	assert.Equal(t, "nil -> 1 -> 9 -> 2 -> 3 -> 4 -> 5 -> nil", forward)
	assert.Equal(t, "nil -> 5 -> 4 -> 3 -> 2 -> 9 -> 1 -> nil", backward)

	forward, backward = strDLL.Show()
	assert.Equal(t, "nil -> Hello! -> I -> am -> just -> testing -> nil", forward)
	assert.Equal(t, "nil -> testing -> just -> am -> I -> Hello! -> nil", backward)
}

func TestIsExistDoublyLL(t *testing.T) {
	numDLL := ds.NewDoublyLinkedList[int]()
	strDLL := ds.NewDoublyLinkedList[string]()

	numDLL.MultiAdd([]int{1, 9, 2, 3, 4, 5})
	strDLL.MultiAdd([]string{"Hello!", "I", "am", "just", "testing"})

	assert.True(t, numDLL.IsExist(4))
	assert.True(t, strDLL.IsExist("am"))

	assert.False(t, numDLL.IsExist(19))
	assert.False(t, strDLL.IsExist("My"))
}

func TestCircularLL(t *testing.T) {
	numCLL := ds.NewCircularLinkedList[int]()
	strCLL := ds.NewCircularLinkedList[string]()

	assert.NotNil(t, numCLL)
	assert.NotNil(t, strCLL)
}

func TestAddCircularLL(t *testing.T) {
	numCLL := ds.NewCircularLinkedList[int]()
	strCLL := ds.NewCircularLinkedList[string]()

	numCLL.Add(19)
	strCLL.Add("Hello")
	assert.Equal(t, "19 -> 19", numCLL.Show())
	assert.Equal(t, "Hello -> Hello", strCLL.Show())
}

func TestMultiAddCircularLL(t *testing.T) {
	numCLL := ds.NewCircularLinkedList[int]()
	strCLL := ds.NewCircularLinkedList[string]()

	numCLL.MultiAdd([]int{1, 9, 2, 3, 4, 5})
	strCLL.MultiAdd([]string{"Hello!", "I", "am", "just", "testing"})
	assert.Equal(t, "1 -> 9 -> 2 -> 3 -> 4 -> 5 -> 1", numCLL.Show())
	assert.Equal(t, "Hello! -> I -> am -> just -> testing -> Hello!", strCLL.Show())
}

func TestIsExistCircularLL(t *testing.T) {
	numCLL := ds.NewCircularLinkedList[int]()
	strCLL := ds.NewCircularLinkedList[string]()

	numCLL.MultiAdd([]int{1, 9, 2, 3, 4, 5})
	strCLL.MultiAdd([]string{"Hello!", "I", "am", "just", "testing"})

	assert.True(t, numCLL.IsExist(4))
	assert.True(t, strCLL.IsExist("am"))

	assert.False(t, numCLL.IsExist(19))
	assert.False(t, strCLL.IsExist("My"))
}
