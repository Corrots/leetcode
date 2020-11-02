package v1

type LRUCache struct {
	head, tail *DLinkedNode
	cache      map[int]*DLinkedNode
	capacity   int
	size       int
}

type DLinkedNode struct {
	key, value int
	prev, next *DLinkedNode
}

func newDLinkedNode(k, v int) *DLinkedNode {
	return &DLinkedNode{key: k, value: v}
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		head:     newDLinkedNode(0, 0),
		tail:     newDLinkedNode(0, 0),
		cache:    make(map[int]*DLinkedNode),
		capacity: capacity,
	}
	// 为链表设置伪头，伪尾
	l.head.next = l.tail
	l.tail.prev = l.head
	return l
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.cache[key]
	if !ok {
		return -1
	}
	// key对应的节点存在cache中
	// 将该节点移动到链表头部
	this.moveToHead(node)
	return node.value
}

func (this *LRUCache) Put(key int, value int) {
	node, ok := this.cache[key]
	if ok {
		// cache中存在key，更新value，并将对应的node移动到链表头
		node.value = value
		this.moveToHead(node)
		return
	}
	// cache未找到key
	// 将{key: node}存入cache
	node = newDLinkedNode(key, value)
	this.cache[key] = node
	// 将节点node加入链表头部
	this.addToHead(node)
	this.size++
	// 若当前缓存数量超过capacity，需删除链表尾部节点，同时删除cache中对应{key: node}，并维护size
	if this.size > this.capacity {
		removed := this.removeTail()
		delete(this.cache, removed.key)
		this.size--
	}
}

// 将指定节点移到链表头部
func (this *LRUCache) moveToHead(node *DLinkedNode) {
	this.removeNode(node)
	this.addToHead(node)
}

// 移除指定节点
func (this *LRUCache) removeNode(node *DLinkedNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

// 将指定节点添加到链表头部
func (this *LRUCache) addToHead(node *DLinkedNode) {
	// 更新node的prev和next
	node.prev = this.head
	node.next = this.head.next
	// 将head重新指向node
	this.head.next.prev = node
	this.head.next = node
}

// 删除链表尾部的节点
func (this *LRUCache) removeTail() *DLinkedNode {
	node := this.tail.prev
	this.removeNode(node)
	return node
}
