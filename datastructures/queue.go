package datastructures

import (
	"fmt"
	"math"
)

var queue []int

// Dequeue - removes from beginning
func dequeue() int {
	if len(queue) > 0 {
		elem := queue[0]
		queue = queue[1:]
		return elem
	}

	return -int(math.NaN()) // :) crap
}

// Enqueue - adds at the end
func enqueue(elem int) {
	queue = append(queue, elem)
}

func TestQueue() {
	enqueue(1)
	enqueue(2)
	enqueue(3)
	enqueue(4)
	fmt.Println(dequeue())
	fmt.Println(dequeue())
	fmt.Println(dequeue())
	fmt.Println(dequeue())
	fmt.Println(dequeue())
}
