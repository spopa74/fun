package various

import "fmt"

/*
Approach 1 - parse from beginning to end, flip between "inSequence" and not, count max len
NOPE - Approach 2 - recurrence == between i and j, if i < i + 1 then max == 1 + max(i+1, j)
											if i >= i + 1 then max == max(i+1, j)
											if i == N then max == 1
NOPE - Approach 3 - recurrence, memoize the i -> j sequences
*/

// returns max length, seq start index
func LongestIncreasing(array []int) (int, int) {

	// return longest1(array)
	return longest2(array, 0, 0)
}

// returns only max, ignores start index
func longest2(array []int, i int, crtMax int) (int, int) {
	if i == len(array)-1 { // last index
		return 1, 0
	}

	if array[i] < array[i+1] {
		newMax, _ := longest2(array, i+1, crtMax)
		newMax += 1
		if newMax > crtMax {
			crtMax = newMax
		}
		return crtMax, 0
	} else {
		newMax, _ := longest2(array, i+1, 0)
		if newMax > crtMax {
			crtMax = newMax
		}
		return crtMax, 0
	}

}

// works, returns both max and start index
func longest1(array []int) (int, int) {

	lastIndex := len(array) - 1

	crtMax := 0
	crtIndex := 0
	newMax := 0
	newIndex := 0
	inSequence := false

	// parse array
	for i := 0; i <= lastIndex; i++ {
		if inSequence {
			if i+1 <= lastIndex { // still going
				if array[i] < array[i+1] {
					fmt.Println("still inSeq, crtMax++, index ", i)
					crtMax++
					continue
				} else { // interrupt sequence
					fmt.Println("interrupt inSeq, index ", i)
					inSequence = false
					if newMax < crtMax {
						fmt.Println("crtMax better, save, index ", i)
						newMax = crtMax
						newIndex = crtIndex
					} else {
						fmt.Println("crtMax NOT better, don't save, index ", i)
						continue
					}
				}
			} else { // at the end, save
				newMax = crtMax
				newIndex = crtIndex
				fmt.Println("at the end inSeq, continue, index ", i)
				continue
			}
		} else { // not inSequence
			if i+1 <= lastIndex { // still going
				if array[i] < array[i+1] { // start sequence
					fmt.Println("found sequence, index ", i)
					inSequence = true
					crtMax = 2
					crtIndex = i
				} else {
					fmt.Println("still not in sequence, continue, index ", i)
					continue // still not inSequence
				}
			} else { // at the end
				fmt.Println("at the end NOT in seq, continue, index ", i)
				continue
			}
		}
	}

	return newMax, newIndex
}
