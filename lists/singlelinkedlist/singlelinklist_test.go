package singlelinkedlist

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSingleLinkedList_Count(t *testing.T) {
	t.Parallel()

	var list SingleLinkedList = NewSingleLinkedList()
	expectedCount := 1

	list.Insert(5)

	actualCount := list.Count()

	assert.Equal(t, expectedCount, actualCount)
}

func TestSingleLinkedList_Find(t *testing.T) {
	t.Parallel()

	var list SingleLinkedList = NewSingleLinkedList()

	list.Insert(1)
	list.Insert(2)
	list.Insert(3)
	list.Insert(4)
	list.Insert(5)

	found := list.Find(3)

	assert.True(t, found)
}

func TestSingleLinkedList_Insert(t *testing.T) {
	t.Parallel()

	var list SingleLinkedList = NewSingleLinkedList()

	expectedValue := 5

	list.Insert(expectedValue)

	assert.Equal(t, expectedValue, list.root.value)
}

func TestSingleLinkedList_Remove(t *testing.T) {
	t.Parallel()

	var list SingleLinkedList = NewSingleLinkedList()

	list.Insert(1)
	list.Insert(2)
	list.Insert(3)
	list.Insert(4)
	list.Insert(5)

	list.Remove(3)

	count := list.Count()

	assert.Equal(t, count, 4)
}