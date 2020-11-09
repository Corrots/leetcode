package main

func main() {

}

// 利用最大堆求解
// 先将数组中前k个元素加入堆中，然后从k索引位置开始向后遍历
// 比较当前元素arr[i]与堆顶元素大小，若arr[i] < 堆顶元素，则推出堆定元素，并将该元素加入堆中
// 重复上述循环，直至循环结束
// 此时堆中存储的即为最小的k个元素

//https://leetcode-cn.com/problems/zui-xiao-de-kge-shu-lcof/
func getLeastNumbers(arr []int, k int) []int {
	var res []int
	n := len(arr)
	if n == 0 || k <= 0 || n < k {
		return res
	}
	pq := &maxHeap{}
	for i := 0; i < k; i++ {
		pq.Push(arr[i])
	}
	for i := k; i < n; i++ {
		if pq.Len() > 0 && arr[i] < pq.Peek().(int) {
			pq.Pop()
			pq.Push(arr[i])
		}
	}

	for i := 0; i < k; i++ {
		res = append(res, pq.Pop().(int))
	}
	return res
}

// 基于go interface实现最大堆
type maxHeap []int

func (h *maxHeap) Len() int {
	return len(*h)
}

func (h *maxHeap) Less(i, j int) bool {
	return (*h)[i] > (*h)[j]
}

func (h *maxHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *maxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *maxHeap) Pop() (v interface{}) {
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return
}

// 查看堆顶元素
func (h *maxHeap) Peek() interface{} {
	return (*h)[len(*h)-1]
}
