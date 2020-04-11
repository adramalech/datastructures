package list

type IList interface {
	Insert(int)
	Count() int
	Find(int) bool
	Remove(int)
}

type IDoubleList interface {
	InsertHead(int)
	InsertTail(int)
	Count() int
	Find(int) bool
	Remove(int)
	FastForward()
	Rewind()
}