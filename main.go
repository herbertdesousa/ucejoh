package main

import (
	"fmt"
)

func main() {
	stack := Stack{}

	stack.Push(NewStackValue(1))

	stack.Push(NewStackValue(2))

	stack.Push(NewStackValue(3))
	stack.Print()
	fmt.Println()
}
