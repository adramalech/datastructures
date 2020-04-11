package binarysearchtree

import (
	"github.com/adramalech/datastructures/models"
	"github.com/adramalech/datastructures/trees"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CountRootNodeShouldBeZero(t *testing.T) {
	t.Parallel()

	// assign
	var root trees.Tree = NewBinarySearchTree()
	expectedCount := 0

	// act
	actualCount := root.Count()

	// assert
	assert.EqualValues(t, expectedCount, actualCount, "the count of single node tree should be one")
}

func Test_RemoveNodeShouldntExist(t *testing.T) {
	t.Parallel()

	var bst trees.Tree = NewBinarySearchTree()

	bst.Insert(5, models.NewPerson("john", "jacob", "smith"))

	bst.Delete(5)

	exists := bst.KeyExists(5)

	assert.False(t, exists)
}

func Test_RemoveNodeFromLargeTree(t *testing.T) {
	t.Parallel()

	var bst trees.Tree = NewBinarySearchTree()

	bst.Insert(10,  models.NewPerson("john", "jacob", "smith"))
	bst.Insert(5,  models.NewPerson("john", "jacob", "smith"))
	bst.Insert(20,  models.NewPerson("john", "jacob", "smith"))
	bst.Insert(15,  models.NewPerson("john", "jacob", "smith"))
	bst.Insert(16,  models.NewPerson("john", "jacob", "smith"))
	bst.Insert(2,  models.NewPerson("john", "jacob", "smith"))
	bst.Insert(8,  models.NewPerson("john", "jacob", "smith"))

	bst.Delete(2)

	exists := bst.KeyExists(2)
	assert.False(t, exists)
}

func Test_NodeInsertedShouldExist(t *testing.T) {
	t.Parallel()

	var bst trees.Tree = NewBinarySearchTree()

	bst.Insert(10, models.NewPerson("test", "tester", "testest"))

	shouldBeTrue := bst.KeyExists(10)

	assert.True(t, shouldBeTrue)
}