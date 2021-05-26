package main

import "container/heap"

func main() {

}

// TopK 问题
//https://leetcode-cn.com/problems/top-k-frequent-elements/submissions/
func topKFrequent(nums []int, k int) []int {
	var res []int
	n := len(nums)
	if n <= 0 || k <= 0 || n < k {
		return res
	}
	// 获取每个元素出现的频率，以val为key，存入map中
	freqMap := make(map[int]int)
	for _, v := range nums {
		freqMap[v]++
	}
	pq := &minHeap{}
	for val, count := range freqMap {
		// 堆中已存入k个元素时
		if pq.Len() == k {
			// 当前元素的频率>堆顶的元素，则替换当前元素
			if count > pq.Peek().(freq).count {
				heap.Pop(pq)
				heap.Push(pq, freq{val: val, count: count})
			}
		} else {
			heap.Push(pq, freq{val: val, count: count})
		}
	}

	for pq.Len() > 0 {
		res = append(res, heap.Pop(pq).(freq).val)
	}
	return res
}

type freq struct {
	val   int
	count int
}

type minHeap []freq

func (h *minHeap) Len() int {
	return len(*h)
}

func (h *minHeap) Less(i, j int) bool {
	return (*h)[i].count < (*h)[j].count
}

func (h *minHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(freq))
}

func (h *minHeap) Pop() interface{} {
	res := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return res
}

func (h *minHeap) Peek() interface{} {
	return (*h)[0]
}
