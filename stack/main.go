package main

import (
	"fmt"
)

func main() {
	var operation int = 0
	var num int
	var stack []int
	var head int = 0
	for {
		fmt.Println("select the operation to perform on stack: ")
		fmt.Scan(&operation)
		switch operation {
		case 1:
			fmt.Println(head)
			Display(stack)

		case 2:
			fmt.Println("Push number to stack: ")
			fmt.Scan(&num)
			stack, head = Push(stack, num)

		case 3:
			var ok bool
			fmt.Println("Pop number from stack")
			stack, head, ok = Pop(stack, head)
			if !ok {
				fmt.Println("stack is empty...")
			}

		case 4:
			peekedElement, ok := Peek(stack, head)
			if ok {

				fmt.Println("Element in head", peekedElement)
			} else {
				fmt.Println("Stack is empty...")
			}
		case 6:
			fmt.Println("head is pointed at: ", stack[head])

		case 7:
			return

		default:
			fmt.Println("select option from 1 to 5")
		}
	}
}
