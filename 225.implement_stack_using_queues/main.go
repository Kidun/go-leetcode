package main

type MyStack struct {
	vals []int
	top  int
}

func Constructor() MyStack {
	return MyStack{vals: make([]int, 0), top: -1}
}

func (this *MyStack) Push(x int) {
	this.top++
	this.vals = append(this.vals, x)
}

func (this *MyStack) Pop() int {
	if this.Empty() {
		return 0
	}
	val := this.vals[this.top]
	this.top--
	this.vals = this.vals[:len(this.vals)-1]
	return val
}

func (this *MyStack) Top() int {
	if this.Empty() {
		return 0
	}
	return this.vals[this.top]
}

func (this *MyStack) Empty() bool {
	return this.top == -1
}

func main() {
	obj := Constructor()
	obj.Push(1)
	param_2 := obj.Pop()
	obj.Push(2)
	param_3 := obj.Top()
	param_4 := obj.Empty()
	println(param_2, " | ", param_3, " | ", param_4)
}
