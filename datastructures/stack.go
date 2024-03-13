package datastructures

import (
	"fmt"
	"math"
)

var stack []int

// Pop - takes from the back
func pop() int {
	if len(stack) > 0 {
		elem := stack[len(stack)-1]
		stack = stack[0 : len(stack)-1]
		return elem
	}

	return -int(math.NaN()) // :) crap
}

// Push - adds at the back
func push(elem int) {
	stack = append(stack, elem)
}

func TestStack() {
	push(1)
	push(2)
	push(3)
	push(4)
	fmt.Println(pop())
	fmt.Println(pop())
	fmt.Println(pop())
	fmt.Println(pop())
	fmt.Println(pop())
}
