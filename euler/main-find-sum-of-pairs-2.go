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
func main_find_sum_of_pairs_2() {

	var numbers = sort.IntSlice([]int{9, 3, 2, 4, 8})
	var givenSum int = 15

	// edge case 1, less than 1 number in slice
	if numbers == nil || numbers.Len() < 2 {
		fmt.Println("too few numbers in slice")
		return
	}

	// sort them
	sort.Sort(numbers)
	fmt.Println(numbers)

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

	found := false
	for i := 0; i <= numbers.Len()-2 && !found; i++ {
		for j := i + 1; j <= numbers.Len()-1 && !found && numbers[i]+numbers[j] <= givenSum; j++ {
			// fmt.Printf("%d, %d, %d\n", numbers[i], numbers[j], numbers[i]+numbers[j])
			if numbers[i]+numbers[j] == givenSum {
				fmt.Printf("Found: %d, %d, %d\n", numbers[i], numbers[j], numbers[i]+numbers[j])
				found = true
				return
			}
		}
	}

	fmt.Println("Not found")
}
