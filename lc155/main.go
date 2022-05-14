package main

/*
设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。

实现 MinStack 类:

MinStack() 初始化堆栈对象。
void push(int val) 将元素val推入堆栈。
void pop() 删除堆栈顶部的元素。
int top() 获取堆栈顶部的元素。
int getMin() 获取堆栈中的最小元素。


来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/min-stack
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

type MinStack struct {
	stackData, stackMin []int
}

func Constructor() MinStack {
	return MinStack{make([]int, 0), make([]int, 0)}
}

func (this *MinStack) Push(val int) {
	this.stackData = append(this.stackData, val)
	if this.isEmptyMin() {
		this.stackMin = append(this.stackMin, val)
	} else {
		min := this.GetMin()
		if val <= min {
			this.stackMin = append(this.stackMin, val)
		}
	}
}

func (this *MinStack) Pop() {
	value := this.stackData[len(this.stackData)-1]
	this.stackData = this.stackData[:len(this.stackData)-1]
	min := this.GetMin()
	if min == value {
		this.stackMin = this.stackMin[:len(this.stackMin)-1]
	}
}

func (this *MinStack) Top() int {
	return this.stackData[len(this.stackData)-1]

}

func (this *MinStack) GetMin() int {
	return this.stackMin[len(this.stackMin)-1]
}

func (this *MinStack) isEmptyMin() bool {
	if len(this.stackMin) == 0 {
		return true
	}
	return false
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
