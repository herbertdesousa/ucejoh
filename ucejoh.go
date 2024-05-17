package main

import "fmt"

type StackValue struct {
	value    int8
	next     *StackValue
	previous *StackValue
}

func NewStackValue(value int8) StackValue {
	stackValue := StackValue{}
	stackValue.value = value
	return stackValue
}

type Stack struct {
	values []StackValue
	head   *StackValue
}

func (stack *Stack) Push(value StackValue) {
	if stack.head == nil {
		stack.head = &value
	} else {
		stack.head.next = &value
		value.previous = stack.head
		stack.head = &value
	}

	stack.values = append(stack.values, value)
}

func printStackValue(value *StackValue) {
	fmt.Printf("Value: %d", value.value)
	fmt.Println()

	if value.previous != nil {
		printStackValue(value.previous)
	}
}

func (stack *Stack) Print() {
	if stack.head == nil {
		fmt.Println("Stack is Empty")
		return
	}

	printStackValue(stack.head)
}
