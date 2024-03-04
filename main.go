package main

import (
	"fmt"
	"fun/leetcode"
)

func main() {

	// meta - reverse grades problem
	// fmt.Println(meta.GetWrongAnswers(0, "ABA"))

	// meta - next line problem
	// fmt.Println(meta.NextLine("111221"))

	// meta - one edit apart problem
	// fmt.Println(meta.CountMinEdits("btsad", 0, "cat", 0))
	// fmt.Println(meta.OneEditApart("cat", "dog"))  // = false
	// fmt.Println(meta.OneEditApart("cat", "cats")) // = true
	// fmt.Println(meta.OneEditApart("cat", "cut"))  // = true
	// fmt.Println(meta.OneEditApart("cat", "cast")) // = true
	// fmt.Println(meta.OneEditApart("cat", "at"))   // = true
	// fmt.Println(meta.OneEditApart("cat", "act"))  // = false

	// meta - additional seats
	// fmt.Println(meta.GetMaxAdditionalDinersCount(10, 1, 2, []int64{2, 6}))
	// fmt.Println(meta.GetMaxAdditionalDinersCount(15, 2, 3, []int64{11, 6, 14}))
	// fmt.Println(meta.GetMaxAdditionalDinersCount(13, 5, 2, []int64{1, 13}))
	// fmt.Println(meta.GetMaxAdditionalDinersCount(13, 0, 3, []int64{5, -11, 13}))
	// fmt.Println(meta.GetMaxAdditionalDinersCount(13, 2, 0, []int64{}))
	// fmt.Println(meta.GetMaxAdditionalDinersCount(13, 15, 2, []int64{1, 2}))
	// fmt.Println(meta.GetMaxAdditionalDinersCount(-13, 15, 2, []int64{1, 2}))
	// fmt.Println(meta.GetMaxAdditionalDinersCount(13, 3, 2, []int64{1, 2, 4}))
	// fmt.Println(meta.GetMaxAdditionalDinersCount(13, 3, 2, []int64{1, 2, 4}))

	// meta - kaitenzushi
	// fmt.Println(meta.GetMaximumEatenDishCount(6, []int32{1, 2, 3, 3, 2, 1}, 1))    // => 5
	// fmt.Println(meta.GetMaximumEatenDishCount(6, []int32{1, 2, 3, 3, 2, 1}, 2))    // => 4
	// fmt.Println(meta.GetMaximumEatenDishCount(7, []int32{1, 2, 1, 2, 1, 2, 1}, 2)) // => 2

	// leetcode - findMedianSortedArrays
	// fmt.Println(leetcode.Median([]int{1, 2, 3}))    // => 2
	// fmt.Println(leetcode.Median([]int{1, 2, 3, 4})) // => 2.5
	fmt.Println(leetcode.FindMedianSortedArrays([]int{1, 2, 3}, []int{1, 2, 3})) // => 2

}
