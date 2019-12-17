package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Stack struct {
	list    []int
	maxSize int
}

func (stack Stack) String() string {
	values := stack.list
	var valuesText []string

	for i := range values {
		number := values[i]
		text := strconv.Itoa(number)
		valuesText = append(valuesText, text)
	}

	result := strings.Join(valuesText, ", ")
	return result
}

func push(stack *Stack, i int) {
	size := len(stack.list)
	if size <= stack.maxSize {
		stack.list = append(stack.list, i)
		fmt.Println(stack.list)
	}
}

func pop(stack *Stack) int {

	l := len(stack.list) - 1
	if l > 0 {
		result := stack.list[l]
		stack.list = stack.list[:l]
		return result
	}

	fmt.Println("List Empty")
	return 0
}

func main() {
	var stack Stack
	stack.maxSize = 10
	fmt.Println(stack)

	push(&stack, 5)
	push(&stack, 6)
	push(&stack, 7)

	fmt.Println(stack)
	fmt.Println(pop(&stack))
	fmt.Println(pop(&stack))
	fmt.Println(pop(&stack))
}
