package binarysearchtree

import "github.com/adramalech/datastructures/models"

type Tree interface {
	Insert(int, models.Person)
	KeyExists(int) bool
	Count() int
	Delete(int)
}
