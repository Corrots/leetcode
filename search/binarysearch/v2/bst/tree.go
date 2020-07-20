package bst

import "fmt"

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

// 前序遍历
func (b *BST) PreOrder() {
	b.root.preOrder()
}

func (n *Node) preOrder() {
	if n != nil {
		fmt.Printf("%d ", n.Key)
		n.left.preOrder()
		n.right.preOrder()
	}
}

// 中序遍历
func (b *BST) InOrder() {
	b.root.inOrder()
}

func (n *Node) inOrder() {
	if n != nil {
		n.left.inOrder()
		fmt.Printf("%d ", n.Key)
		n.right.inOrder()
	}
}

// 后序遍历
func (b *BST) PostOrder() {
	b.root.postOrder()
}

func (n *Node) postOrder() {
	if n != nil {
		n.left.postOrder()
		n.right.postOrder()
		fmt.Printf("%d ", n.Key)
	}
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
