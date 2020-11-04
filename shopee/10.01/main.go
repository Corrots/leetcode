package main

func main() {

}

// 思路：双指针

//https://leetcode-cn.com/problems/sorted-merge-lcci/
func merge(A []int, m int, B []int, n int) {
	p, q := m-1, n-1

	var val int
	for tail := m + n - 1; p >= 0 || q >= 0; tail-- {
		if p < 0 {
			val = B[q]
			q--
		} else if q < 0 {
			val = A[p]
			p--
		} else if A[p] > B[q] {
			val = A[p]
			p--
		} else {
			val = B[q]
			q--
		}
		A[tail] = val
	}
}
