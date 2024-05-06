package tree

import "fmt"

/*
tree stack
*/

type QueueStack struct {
	//data
	data []interface{}
	//stack
	top int
}

func NewQueueStack() *QueueStack {
	return &QueueStack{
		data: make([]interface{}, 0, 32),
		top:  -1,
	}
}

func (this *QueueStack) IsEmpty() bool {
	if this.top < 0 {
		return true
	}
	return false
}

func (this *QueueStack) Push(v interface{}) {
	if this.top < 0 {
		this.top = 0
	} else {
		this.top += 1
	}

	if this.top > len(this.data)-1 {
		this.data = append(this.data, v)
	} else {
		this.data[this.top] = v
	}
}

func (this *QueueStack) Pop() interface{} {
	if this.IsEmpty() {
		return nil
	}
	v := this.data[0]
	this.data = this.data[1:]
	this.top -= 1
	return v
}

func (this *QueueStack) Head() interface{} {
	if this.IsEmpty() {
		return nil
	}
	return this.data[0]
}

func (this *QueueStack) Top() interface{} {
	if this.IsEmpty() {
		return nil
	}
	return this.data[this.top]
}

func (this *QueueStack) Flush() {
	this.top = -1
}

func (this *QueueStack) Print() {
	if this.IsEmpty() {
		fmt.Println("empty statck")
	} else {
		for i := this.top; i >= 0; i-- {
			fmt.Println(this.data[i])
		}
	}
}
