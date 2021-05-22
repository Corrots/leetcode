package main

/**
一个栈作为输入栈，用于push数据；一个栈作为输出栈，用于pop和peek
当输出栈为空时，将输入栈中的数据全部取出放入输出栈中
*/

func main() {

}

type MyQueue struct {
	inStack  []int
	outStack []int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{}
}

/** Push element x to the back of queue. */
func (q *MyQueue) Push(x int) {
	q.inStack = append(q.inStack, x)
}

// 将输入栈中的数据转入输出栈中
func (q *MyQueue) in2out() {
	for len(q.inStack) > 0 {
		q.outStack = append(q.outStack, q.inStack[len(q.inStack)-1])
		q.inStack = q.inStack[:len(q.inStack)-1]
	}
}

/** Removes the element from in front of queue and returns that element. */
func (q *MyQueue) Pop() int {
	if len(q.outStack) == 0 {
		q.in2out()
	}
	out := q.outStack[len(q.outStack)-1]
	q.outStack = q.outStack[:len(q.outStack)-1]
	return out
}

/** Get the front element. */
func (q *MyQueue) Peek() int {
	if len(q.outStack) == 0 {
		q.in2out()
	}
	return q.outStack[len(q.outStack)-1]
}

/** Returns whether the queue is empty. */
func (q *MyQueue) Empty() bool {
	return len(q.inStack) == 0 && len(q.outStack) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
