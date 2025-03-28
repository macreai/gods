package test

import (
	"testing"

	"github.com/macreai/gods/ds"
	"github.com/stretchr/testify/assert"
)

func TestTree(t *testing.T) {
	trees := ds.NewBinaryTree[int]()
	assert.NotNil(t, trees)
}

func TestAddTree(t *testing.T) {
	tree := ds.NewBinaryTree[int]()
	tree.Add(10)
	assert.True(t, tree.IsExist(10))
}

func TestMultiAddTree(t *testing.T) {
	tree := ds.NewBinaryTree[int]()
	values := []int{1, 2, 4, 5, 6, 7, 8, 23, 54, 62, 12, 86, 92, 54, 39, 47}
	tree.MultiAddd(values)
	for _, v := range values {
		assert.True(t, tree.IsExist(v))
	}
}
