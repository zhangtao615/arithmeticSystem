package createStack

import "fmt"

type StackMethods interface {
	Push(val interface{})
	Pop() interface{}
	Peek() interface{}
	GetSize() int
	IsEmpty() bool
}

type Stack struct {
	stack []interface{}
	size  int
	count int
}

func (stack *Stack) Push(val interface{}) {
	if stack.size == stack.count {
		fmt.Println("栈空间已满")
		return
	} else {
		stack.stack = append(stack.stack, val)
	}
}

func (stack *Stack) Pop() interface{} {
	if stack.IsEmpty() {
		fmt.Println("栈为空")
		return nil
	} else {
		return stack.stack[len(stack.stack)-1 : len(stack.stack)]
	}
}

func (stack *Stack) Peek() interface{} {
	if stack.IsEmpty() {
		fmt.Println("栈为空")
		return nil
	} else {
		return stack.stack[0]
	}
}

func (stack *Stack) GetSize() int {
	return len(stack.stack)
}

func (stack *Stack) IsEmpty() bool {
	return len(stack.stack) == 0
}

func NewStack(size int) StackMethods {
	return &Stack{
		stack: make([]interface{}, size),
		count: 0,
		size:  size,
	}
}
