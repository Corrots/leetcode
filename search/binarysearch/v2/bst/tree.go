package bst

import "fmt"

type BST struct {
	root  *Node
	count int
}

func NewBST() *BST {
	return &BST{
		root:  nil,
		count: 0,
	}
}

func (b *BST) RemoveMin() {
	if b.root != nil {
		b.root = b.root.removeMin()
		b.count--
	}
}

func (b *BST) RemoveMax() {
	if b.root != nil {
		b.root = b.root.removeMax()
		b.count--
	}
}

// 返回二叉搜索树中的最小键值
func (b *BST) Minimum() int {
	return b.root.minimum().Key
}

func (b *BST) Maximum() int {
	return b.root.maximum().Key
}

// 前序遍历
func (b *BST) PreOrder() {
	b.root.preOrder()
}

// 中序遍历
func (b *BST) InOrder() {
	b.root.inOrder()
}

// 后序遍历
func (b *BST) PostOrder() {
	b.root.postOrder()
}

// 层序遍历(广度优先遍历)
func (b *BST) levelOrder() {
	var q []*Node
	q = append(q, b.root)
	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		fmt.Printf("%d ", node.Key)
		if node.left != nil {
			q = append(q, node.left)
		}
		if node.right != nil {
			q = append(q, node.right)
		}
	}
}

func (b *BST) Search(key int) interface{} {
	return b.root.search(key)
}

func (b *BST) Contain(key int) bool {
	return b.root.contain(key)
}

func (b *BST) Insert(key int, value interface{}) {
	b.root = b.root.insert(key, value)
}

func (b *BST) Len() int {
	return b.count
}

func (b *BST) IsEmpty() bool {
	return b.count == 0
}

func (b *BST) Destroy() {
	b.root.destroy()
	b.count--
}
