package singlelinkedlist

import (
	list "github.com/adramalech/datastructures/lists"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSingleLinkedList_Count(t *testing.T) {
	t.Parallel()

	var l list.IList = NewSingleLinkedList()

	expectedCount := 1

	l.Insert(5)

	actualCount := l.Count()

	assert.Equal(t, expectedCount, actualCount)
}

func TestSingleLinkedList_Find(t *testing.T) {
	t.Parallel()

	var l list.IList = NewSingleLinkedList()

	l.Insert(1)
	l.Insert(2)
	l.Insert(3)
	l.Insert(4)
	l.Insert(5)

	found := l.Find(3)

	assert.True(t, found)
}

func TestSingleLinkedList_Insert(t *testing.T) {
	t.Parallel()

	var l list.IList = NewSingleLinkedList()

	expectedValue := 5

	l.Insert(expectedValue)

	shouldBeFound := l.Find(5)
	assert.True(t, shouldBeFound)
}

func TestSingleLinkedList_Remove(t *testing.T) {
	t.Parallel()

	var l list.IList = NewSingleLinkedList()

	l.Insert(1)
	l.Insert(2)
	l.Insert(3)
	l.Insert(4)
	l.Insert(5)

	l.Remove(3)

	count := l.Count()

	assert.Equal(t, count, 4)
}
