package main

import "fmt"

func main() {

	// reverse grades problem
	// fmt.Println(getWrongAnswers(0, "ABA"))

	// next line problem
	// fmt.Println(nextLine("111221"))

	// one edit apart problem
	// fmt.Println(countMinEdits("btsad", 0, "cat", 0))
	// fmt.Println(oneEditApart("cat", "dog"))  // = false
	// fmt.Println(oneEditApart("cat", "cats")) // = true
	// fmt.Println(oneEditApart("cat", "cut"))  // = true
	// fmt.Println(oneEditApart("cat", "cast")) // = true
	// fmt.Println(oneEditApart("cat", "at"))   // = true
	// fmt.Println(oneEditApart("cat", "act"))  // = false

	// additional seats
	// fmt.Println(getMaxAdditionalDinersCount(10, 1, 2, []int64{2, 6}))
	// fmt.Println(getMaxAdditionalDinersCount(15, 2, 3, []int64{11, 6, 14}))
	// fmt.Println(getMaxAdditionalDinersCount(13, 5, 2, []int64{1, 13}))
	// fmt.Println(getMaxAdditionalDinersCount(13, 0, 3, []int64{5, -11, 13}))
	// fmt.Println(getMaxAdditionalDinersCount(13, 2, 0, []int64{}))
	// fmt.Println(getMaxAdditionalDinersCount(13, 15, 2, []int64{1, 2}))
	// fmt.Println(getMaxAdditionalDinersCount(-13, 15, 2, []int64{1, 2}))
	// fmt.Println(getMaxAdditionalDinersCount(13, 3, 2, []int64{1, 2, 4}))
	// fmt.Println(getMaxAdditionalDinersCount(13, 3, 2, []int64{1, 2, 4}))

	// kaitenzushi
	fmt.Println(getMaximumEatenDishCount(6, []int32{1, 2, 3, 3, 2, 1}, 1)) // => 5
	// fmt.Println(getMaximumEatenDishCount(6, []int32{1, 2, 3, 3, 2, 1}, 2)) // => 4
	// fmt.Println(getMaximumEatenDishCount(7, []int32{1, 2, 1, 2, 1, 2, 1}, 2)) // => 2
}
