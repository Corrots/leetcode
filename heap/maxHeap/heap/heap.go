package heap

import "fmt"

// 最大堆
type Heap struct {
	data     []int
	count    int
	capacity int
}

func New(capacity int) *Heap {
	return &Heap{
		data:     make([]int, capacity+1, capacity+1),
		capacity: capacity,
	}
}

func Heapify(nums []int, n int) *Heap {
	var maxHeap Heap
	maxHeap.data = make([]int, n+1)
	maxHeap.capacity = n
	for i := 0; i < n; i++ {
		maxHeap.data[i+1] = nums[i]
	}
	maxHeap.count = n
	for i := maxHeap.count / 2; i >= 1; i-- {
		maxHeap.shiftDown(i)
	}
	return &maxHeap
}

func (h *Heap) String() string {
	res := "["
	for i := 1; i <= h.count; i++ {
		res += fmt.Sprintf("%d ", h.data[i])
	}
	res += "]"
	return res
}

func (h *Heap) Len() int {
	return h.count
}

func (h *Heap) IsEmpty() bool {
	return h.count == 0
}

func (h *Heap) Insert(n int) {
	if h.count+1 > h.capacity {
		arr := make([]int, h.capacity*2, h.capacity*2)
		copy(arr, h.data)
	}
	h.data[h.count+1] = n
	h.count++
	h.shiftUp(h.count)
}

func (h *Heap) shiftUp(k int) {
	for k > 1 && h.data[k/2] < h.data[k] {
		h.data[k], h.data[k/2] = h.data[k/2], h.data[k]
		k /= 2
	}
}

// 推出堆中最大值
func (h *Heap) ExtractMax() int {
	if h.count == 0 {
		return 0
	}
	max := h.data[1]
	// 将根节点元素与最后元素交换位置
	h.data[1] = h.data[h.count]
	h.count--
	// 调整第一个元素在堆中的位置
	h.shiftDown(1)
	return max
}

func (h *Heap) shiftDown(k int) {
	// 只要有子树就一直遍历下去
	for 2*k <= h.count {
		// 此轮循环h.data[k]和h.data[j]互换位置
		j := 2 * k
		// 如果右子树存在，且右子树>左子树，变更为右子树
		if j+1 <= h.count && h.data[j+1] > h.data[j] {
			j += 1
		}
		// 循环终止条件
		if h.data[k] >= h.data[j] {
			break
		}
		// 父节点与子树的元素互换位置
		h.data[k], h.data[j] = h.data[j], h.data[k]
		k = j
	}
}
