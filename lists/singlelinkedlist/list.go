package singlelinkedlist

type List interface {
	Insert(int)
	Count() int
	Find(int) bool
	Remove(int)
}