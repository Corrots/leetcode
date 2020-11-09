package main

import "container/heap"

func main() {

}

// 用哈希表记录每个元素出现的次数
// 构建使用最小堆，遍历哈希表，当最小堆中的元素数量<k时，将当前数据对加入堆中
// 若堆中元素数量==k时，比较堆顶元素的count与哈希表中count的大小
// 若<，推出堆顶元素，并将该元素加入堆中，遍历循环至结束
// 最后取出堆中所有元素加入结果集

//https://leetcode-cn.com/problems/top-k-frequent-elements/
func topKFrequent(nums []int, k int) []int {
	var res []int
	n := len(nums)
	if n == 0 || k <= 0 || n < k {
		return res
	}
	freqMap := make(map[int]int)
	for _, v := range nums {
		freqMap[v]++
	}

	pq := &minHeap{}
	// 注意遍历时的变量定义,变量k已作为参数传入
	for val, count := range freqMap {
		if pq.Len() == k {
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

func (h *minHeap) Pop() (v interface{}) {
	lastIdx := len(*h) - 1
	*h, v = (*h)[:lastIdx], (*h)[lastIdx]
	return
}

func (h *minHeap) Peek() interface{} {
	return (*h)[0]
}
