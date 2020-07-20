package bst

type BST struct {
	root  *Node
	count int
}

type Node struct {
	Key   int
	Value interface{}
	left  *Node
	right *Node
}

func (b *BST) Search(key int) interface{} {
	return b.root.search(key)
}

func (n *Node) search(key int) interface{} {
	if n == nil {
		return nil
	}
	if key == n.Key {
		return n.Value
	} else if key < n.Key {
		return n.left.search(key)
	} else {
		return n.right.search(key)
	}
}

func (b *BST) Contain(key int) bool {
	return b.root.contain(key)
}

func (n *Node) contain(key int) bool {
	if n == nil {
		return false
	}

	if key == n.Key {
		return true
	} else if key < n.Key {
		return n.left.contain(key)
	} else {
		return n.right.contain(key)
	}
}

func (b *BST) Insert(key int, value interface{}) {
	b.root = b.root.insert(key, value)
}

func (n *Node) insert(key int, value interface{}) *Node {
	if n == nil {
		return newNode(key, value)
	}

	if key == n.Key {
		n.Value = value
	} else if key < n.Key {
		n.left = n.left.insert(key, value)
	} else {
		n.right = n.right.insert(key, value)
	}
	return n
}

func NewBST() *BST {
	return &BST{
		root:  nil,
		count: 0,
	}
}

func (b *BST) Len() int {
	return b.count
}

func (b *BST) IsEmpty() bool {
	return b.count == 0
}

func newNode(key int, value interface{}) *Node {
	return &Node{
		Key:   key,
		Value: value,
	}
}
