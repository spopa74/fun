package google

import (
	"errors"
	"fmt"
)

type node struct {
	previous  *node
	wheresMax *node
	value     int
}

var myStack *node

/*
Pop element:
- return last int if there's an element
- the stack should have less one element
- return error if stack is empty
*/
func (s *node) pop() (int, *node, error) {
	if s == nil {
		return 0, nil, errors.New("stack empty")
	}

	return s.value, s.previous, nil
}

/*
Put element:
- create a new element with the given value
- create a new 'node'
*/
func (s *node) put(element int) *node {
	newNode := new(node)
	newNode.value = element
	newNode.previous = s

	if s == nil || element >= s.wheresMax.value {
		newNode.wheresMax = newNode
	} else {
		// find the spot
		newNode.wheresMax = s.wheresMax
	}

	return newNode
}

func main_max_stack() {

	values := []int{1, 4, 2, 88, 33, 44, 5, 6}
	for i := 0; i < len(values); i++ {
		myStack = myStack.put(values[i])
		fmt.Printf("Putting value %d\n", values[i])
	}

	for i := 0; i < 15; i++ {
		value, node, err := myStack.pop()
		if err == nil {
			fmt.Printf("Found element %d, max on stack is %d\n", value, myStack.wheresMax.value)
		} else {
			fmt.Printf("error %s\n", err.Error())
		}

		myStack = node
	}

	fmt.Println()
}
