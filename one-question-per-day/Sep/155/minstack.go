package minstack

import "math"

//当一个元素要入栈时，我们取当前辅助栈的栈顶存储的最小值，与当前元素比较得出最小值，将这个最小值插入辅助栈中；
//当一个元素要出栈时，我们把辅助栈的栈顶元素也一并弹出；
//在任意一个时刻，栈内元素的最小值就存储在辅助栈的栈顶元素中。

type MinStack struct {
	stack    []int
	minStack []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		stack:    []int{},
		minStack: []int{math.MaxInt64},
	}
}

// 向stack中插入元素时，同步维护minStack
func (this *MinStack) Push(x int) {
	this.stack = append(this.stack, x)
	// 获取minStack当前栈顶的元素(即当前最小值)
	top := this.minStack[len(this.minStack)-1]
	// 比较minStack当前栈顶的元素和x,将较小值加入minStack，其即为当前minStack的最小值
	this.minStack = append(this.minStack, min(top, x))
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
	this.minStack = this.minStack[:len(this.minStack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack)-1]
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
