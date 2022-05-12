package main

import "container/heap"

func main() {

}

// TopK 问题
//https://leetcode-cn.com/problems/top-k-frequent-elements/submissions/
func topKFrequent(nums []int, k int) []int {
	var res []int
	if len(nums) == 0 || k <= 0 || len(nums) < k {
		return res
	}
	// 遍历统计所有元素出现的次数
	freMap := make(map[int]int)
	for _, val := range nums {
		freMap[val]++
	}

	pq := &minHeap{}
	for k, count := range freMap {
		if pq.Len() == k {
			if count > pq.Peek().(freq).count {
				heap.Pop(pq)
				heap.Push(pq, freq{val: k, count: count})
			}
		} else {
			heap.Push(pq, freq{val: k, count: count})
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

func (m *minHeap) Len() int {
	return len(*m)
}

func (m *minHeap) Less(i, j int) bool {
	// 最小堆
	return (*m)[i].count < (*m)[j].count
}

func (m *minHeap) Swap(i, j int) {
	(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
}

func (m *minHeap) Push(x interface{}) {
	*m = append(*m, x.(freq))
}

func (m *minHeap) Pop() interface{} {
	res := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return res
}

func (m *minHeap) Peek() interface{} {
	return (*m)[0]
}
