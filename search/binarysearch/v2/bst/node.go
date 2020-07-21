package bst

import "fmt"

type Node struct {
	Key   int
	Value interface{}
	left  *Node
	right *Node
}

func newNode(key int, value interface{}) *Node {
	return &Node{
		Key:   key,
		Value: value,
	}
}

// 删除以node为根的二分搜索树中的最小节点
func (n *Node) removeMin() *Node {
	if n.left == nil {
		// @TODO 删除内存空间，count--
		return n.right
	}
	// 继续寻找最小值
	n.left = n.left.removeMin()
	return n
}

func (n *Node) removeMax() *Node {
	if n.right == nil {
		// @TODO 删除内存空间，count--
		return n.left
	}
	// 继续寻找最大值
	n.right = n.right.removeMax()
	return n
}

// 在以node为根的二叉搜索树中，返回最小键值的节点
func (n *Node) minimum() *Node {
	if n.left == nil {
		return n
	}
	return n.left.minimum()
}

func (n *Node) maximum() *Node {
	if n.right == nil {
		return n
	}
	return n.right.maximum()
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

// 对Node进行前序遍历
func (n *Node) preOrder() {
	if n != nil {
		fmt.Printf("%d ", n.Key)
		n.left.preOrder()
		n.right.preOrder()
	}
}

// 对Node进行中序遍历
func (n *Node) inOrder() {
	if n != nil {
		n.left.inOrder()
		fmt.Printf("%d ", n.Key)
		n.right.inOrder()
	}
}

// 对Node进行后序遍历
func (n *Node) postOrder() {
	if n != nil {
		n.left.postOrder()
		n.right.postOrder()
		fmt.Printf("%d ", n.Key)
	}
}

// 析构函数，对二叉搜索树进行空间释放
func (n *Node) destroy() {
	if n != nil {
		n.left.destroy()
		n.right.destroy()
		n.destroy()
	}
}
