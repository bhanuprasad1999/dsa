package main

import "fmt"

func Push(stack []int, num int) ([]int, int) {
	fmt.Println("push operation")
	stack = append(stack, num)
	head := len(stack) - 1
	return stack, head
}

func Pop(stack []int, head int) ([]int, int, bool) {
	if len(stack) == 0 {
		return stack, head, false
	}
	fmt.Println("popped number", stack[head])
	stack = stack[:head]
	head = head - 1
	return stack, head, true
}

func IsEmpty(stack []int) bool {
	fmt.Println("isempty operation")
	return len(stack) == 0
}

func Peek(stack []int, head int) (int, bool) {
	fmt.Println("peek operation")
	if len(stack) == 0 {
		return 0, false
	} else {

		return stack[head], true
	}
}

func Display(stack []int) {
	fmt.Println("display stack details")
	fmt.Println(stack)
}
