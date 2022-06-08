package main

import (
	"algorithmSystem/Chapter1/Stack/createStack"
	"fmt"
)

func main() {
	//stack := createStack.NewStack(10)
	//stack.Push(1)
	//stack.Push(2)
	//stack.Push(3)
	//fmt.Println(stack)

	isValid := createStack.IsValid("[][](){}")
	fmt.Println(isValid)
}
