package main

import (
	"fmt"
	"sort"
)

/* find pairs which sum to some value
Questions:
- sorted or unsorted?
- stop at first pair or find all?
- edge cases?

Initial discussion:
- naive approach == n x n checks => try, if succeeds, exit
- less naive == sort then try
	- sort == n x log(n)
	- checks == fewer than n x n

*/

// EASY

// define a pair struct
type pair struct {
	p1 int
	p2 int
}

// main func
func main_sum_of_pairs() {

	var numbers = sort.IntSlice([]int{9, 3, 2, 4, 8})
	var givenSum int = 11

	// edge case 1, less than 1 number in slice
	if numbers == nil || numbers.Len() < 2 {
		fmt.Println("too few numbers in slice")
		return
	}

	// sort them
	sort.Sort(numbers)

	// edge case 2, sum of 2 min > given sum
	if numbers[0]+numbers[1] > givenSum {
		fmt.Println("sum of smallest greater than given sum, exit")
		return
	}

	// edge case 3, sum of last 2 < given sum
	if numbers[numbers.Len()-1]+numbers[numbers.Len()-2] < givenSum {
		fmt.Println("sum of greatest smaller than given sum, exit")
		return
	}

	// there's a chance
	pair := findPairForSum(numbers, 0, 1, numbers.Len()-1, givenSum)
	if pair == nil {
		fmt.Println("didn't find, exit")
		return
	}

	fmt.Println(pair)
}

// won't check
func findPairForSum(slice []int, current int, start int, end int, sum int) *pair {

	if start >= len(slice) || slice[current]+slice[start] > sum {
		return nil
	}

	fmt.Println("current: %i, start: %i, end: %, this sum: %i", current, start, end, slice[current]+slice[start])

	// yey!
	if slice[current]+slice[start] == sum {
		return &pair{
			p1: slice[current],
			p2: slice[start],
		}
	}

	// recurse 1
	found := findPairForSum(slice, current, start+1, end, sum)
	if found != nil {
		return found
	}

	// recurse 2
	return findPairForSum(slice, current+1, start+1, end, sum)

}
