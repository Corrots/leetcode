package main

type LRUCache struct {
	head, tail *listNode // 伪头，伪尾
	cache      map[int]*listNode
	capacity   int
	size       int
}

type listNode struct {
	key, val   int
	prev, next *listNode
}

func NewListNode(key, val int) *listNode {
	return &listNode{key: key, val: val}
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		head:     NewListNode(0, 0),
		tail:     NewListNode(0, 0),
		cache:    make(map[int]*listNode),
		capacity: capacity,
		size:     0,
	}
	l.head.next = l.tail
	l.tail.prev = l.head
	return l
}

func (l *LRUCache) Get(key int) int {
	node, existed := l.cache[key]
	if !existed {
		return -1
	}
	// 存在，将node移动到链表头部
	l.moveToHead(node)
	return node.val
}

func (l *LRUCache) Put(key int, value int) {
	node, existed := l.cache[key]
	if existed {
		// 存在，更新val，并将node移动到链表头部
		node.val = value
		l.moveToHead(node)
		return
	}
	// 不存在
	node = NewListNode(key, value)
	l.cache[key] = node
	l.addToHead(node)
	l.size++
	if l.size > l.capacity {
		// 移除链表尾部节点
		tail := l.removeTail()
		delete(l.cache, tail.key)
		l.size--
	}
}

// 2 steps:
//	removeNode
//	addToHead
func (l *LRUCache) moveToHead(node *listNode) {
	l.removeNode(node)
	l.addToHead(node)
}

func (l *LRUCache) removeNode(node *listNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (l *LRUCache) addToHead(node *listNode) {
	node.prev = l.head
	node.next = l.head.next
	// 更新头节点
	l.head.next.prev = node
	l.head.next = node
}

func (l *LRUCache) removeTail() *listNode {
	node := l.tail.prev
	l.removeNode(node)
	return node
}
