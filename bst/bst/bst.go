package bst

type Node struct {
	e     interface{}
	left  *Node
	right *Node
}

type BST struct {
	root *Node
	size int
}

func newNode(e interface{}) *Node {
	return &Node{e: e}
}

func New() *BST {
	return &BST{}
}

func (b *BST) Len() int {
	return b.size
}

func (b *BST) IsEmpty() bool {
	return b.size == 0
}
