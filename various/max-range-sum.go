package various

import "fmt"

var maxR []int

func MaxRangeSum(array []int) {
	maxRange(array, len(array)-1)
	fmt.Println(maxR)
}

func maxRange(array []int, n int) int {
	if n == 0 {
		if array[0] > 0 {
			maxR = append(maxR, array[0])
			return array[0]
		} else {
			maxR = append(maxR, 0)
			return 0
		}
	}

	recurrence := maxRange(array, n-1) + array[n]
	if recurrence > 0 {
		maxR = append(maxR, recurrence)
		return recurrence
	} else {
		maxR = append(maxR, 0)
		return 0
	}

}
