package main

import "container/heap"

func main() {

}

// 记录每个点到原点距离的平方sqrt,并将sqrt与point对应记录
// 遍历记录的同时构建最大堆，当堆中元素数量<K时，继续向堆中推入元素
// 当堆中len==k时，判断堆顶元素与当前遍历元素的大小

//https://leetcode-cn.com/problems/k-closest-points-to-origin/
func kClosest(points [][]int, k int) [][]int {
	var res [][]int
	n := len(points)
	if n == 0 || k <= 0 || n < k {
		return res
	}
	// 将前k个元素加入堆中
	pq := make(maxHeap, k)
	for i, p := range points[:k] {
		pq[i] = pair{point: p, dist: p[0]*p[0] + p[1]*p[1]}
	}
	heap.Init(&pq)

	for _, p := range points[k:] {
		if dist := p[0]*p[0] + p[1]*p[1]; dist < pq.Peek().(pair).dist {
			//heap.Pop(&pq)
			//heap.Push(&pq, pair{point: p, dist: dist})
			pq[0] = pair{p, dist}
			heap.Fix(&pq, 0)
		}
	}

	for pq.Len() > 0 {
		res = append(res, heap.Pop(&pq).(pair).point)
	}
	return res
}

type pair struct {
	point []int
	dist  int
}

type maxHeap []pair

func (h *maxHeap) Len() int {
	return len(*h)
}

func (h *maxHeap) Less(i, j int) bool {
	return (*h)[i].dist > (*h)[j].dist
}

func (h *maxHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *maxHeap) Push(x interface{}) {
	*h = append(*h, x.(pair))
}

func (h *maxHeap) Pop() (v interface{}) {
	lastIndex := len(*h) - 1
	*h, v = (*h)[:lastIndex], (*h)[lastIndex]
	return
}
func (h *maxHeap) Peek() interface{} {
	return (*h)[0]
}
