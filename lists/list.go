package list

type IList interface {
	Insert(int)
	Count() int
	Find(int) bool
	Remove(int)
}
