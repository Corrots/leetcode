package set

type Set interface {
	Add(interface{})
	Remove(interface{})
	Contains(interface{})
	Len() int
	IsEmpty() bool
}
