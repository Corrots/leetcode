package main

import "container/heap"

func main() {

}

// 求数组中最小的k个数
// 利用最大堆求解
// 先将数组中前k个元素加入堆中，然后从k索引位置开始向后遍历
// 比较当前元素arr[i]与堆顶元素大小，若arr[i] < 堆顶元素，则推出堆顶元素，并将该元素加入堆中
// 重复上述循环，直至循环结束
// 此时堆中存储的即为最小的k个元素

//https://leetcode-cn.com/problems/zui-xiao-de-kge-shu-lcof/
func getLeastNumbers(arr []int, k int) []int {
	var res []int
	n := len(arr)
	if n <= 0 || k <= 0 || n < k {
		return res
	}
	pq := &MaxHeap{}
	// 先将数组的前k个元素加入堆中
	for i := 0; i < k; i++ {
		heap.Push(pq, arr[i])
	}
	heap.Init(pq)
	// 继续遍历数组
	for i := k; i < n; i++ {
		// 每个元素与堆顶的元素比较大小，若<堆顶元素，则推出堆顶元素，将当前元素加入堆中
		if pq.Len() > 0 && arr[i] < pq.Peek() {
			heap.Pop(pq)
			heap.Push(pq, arr[i])
		}
	}

	for i := 0; i < k; i++ {
		res = append(res, heap.Pop(pq).(int))
	}
	return res
}

type MaxHeap []int

func (h *MaxHeap) Len() int {
	return len(*h)
}

// Less 最大堆
func (h *MaxHeap) Less(i, j int) bool {
	return (*h)[i] > (*h)[j]
}

func (h *MaxHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	res := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return res
}

func (h *MaxHeap) Peek() int {
	return (*h)[0]
}
