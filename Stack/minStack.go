package Stack

import "fmt"

type MinStack struct {
	top *StackObj
	min int
}

type StackObj struct {
	data    int
	next    *StackObj
	prevMin int
}

var myStack MinStack = MinStack{top: nil}

func Constructor() MinStack {
	return myStack
}

// func (this *MinStack) Print() {
// 	itrObj := this.top
// 	for itrObj != nil {
// 		fmt.Println(itrObj.data)
// 		itrObj = itrObj.next
// 	}
// }

func (this *MinStack) Push(val int) {
	newObj := &StackObj{}
	if this.top == nil {
		newObj.data = val
		newObj.next = nil
		newObj.prevMin = val
		this.min = val
	} else {
		newObj.data = val
		// This will store the previous min so that if we pop the top obj,
		// there will be a previous min to access in O(1) time
		newObj.prevMin = this.min
		//This is a stack so next will be the last topmost element
		newObj.next = this.top
	}
	this.top = newObj
	if this.min > newObj.data {
		this.min = newObj.data
	}
}

func (this *MinStack) Pop() {
	// if this is the last object
	if this.top.next == nil {
		this.top = nil
		return
	}
	tempObj := this.top
	this.min = this.top.prevMin
	this.top = this.top.next
	fmt.Println("Popped : ", tempObj.data)
}

func (this *MinStack) Top() int {
	return this.top.data
}

func (this *MinStack) GetMin() int {
	return this.min
}
