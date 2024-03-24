package datastructures

import "fmt"

var array = []int{1, 2, 3, 4}

func createPermutations(input []int) [][]int {
	var ret [][]int = make([][]int, 0)

	for i := 0; i < len(input); i++ {
		// find out rest
		var cp []int = make([]int, len(input))
		copy(cp, input)
		startElem := cp[i]
		cp = append(cp[:i], cp[i+1:]...)
		// if one elem left, no recurrent calls
		if len(cp) == 1 {
			ret = append(ret, []int{startElem, cp[0]})
		} else {
			smallerPerms := createPermutations(cp)
			for j := 0; j < len(smallerPerms); j++ {
				// fmt.Println(smallerPerms[j])
				ret = append(ret, append([]int{startElem}, smallerPerms[j]...))
			}
		}
	}
	return ret
}

func TestPerm() {
	fmt.Println(createPermutations(array))
}
