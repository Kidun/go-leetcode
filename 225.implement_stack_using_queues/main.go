package main

// single queue solution
type MyStack struct {
	q []int
}

func Constructor() MyStack {
	return MyStack{}
}

func (this *MyStack) Push(x int) {
	this.q = append(this.q, x)
	for i := 0; i < len(this.q)-1; i++ {
		this.q = append(this.q, this.q[0])
		this.q = this.q[1:]
	}
}

func (this *MyStack) Pop() int {
	val := this.q[0]
	this.q = this.q[1:]
	return val
}

func (this *MyStack) Top() int {
	return this.q[0]
}

func (this *MyStack) Empty() bool {
	return len(this.q) == 0
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
