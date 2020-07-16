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

func (h *Heap) String() string {
	res := "["
	for i := 1; i <= h.count; i++ {
		res += fmt.Sprintf("%d ", h.data[i])
	}
	res += "]"
	return res
}
