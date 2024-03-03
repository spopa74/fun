package meta

// WARMUP
func oneEditApart(s1 string, s2 string) bool {

	return countMinEdits(s1, 0, s2, 0) == 1

}

func countMinEdits(s1 string, i1 int, s2 string, i2 int) int {

	if s1 == "" || i1 == len(s1) {
		// fmt.Println("s1 empty")
		return len(s2) - i2
	}

	if s2 == "" || i2 == len(s2) {
		// fmt.Println("s2 empty")
		return len(s1) - i1
	}

	// change letter in place, count rest
	clip := 0
	if s1[i1] == s2[i2] {
		clip = countMinEdits(s1, i1+1, s2, i2+1)
	} else {
		clip = 1 + countMinEdits(s1, i1+1, s2, i2+1)
	}

	// remove letter s1, count rest
	rl1 := 1 + countMinEdits(s1, i1+1, s2, i2)

	// remove letter s2, count rest
	rl2 := 1 + countMinEdits(s1, i1, s2, i2+1)

	return min(clip, rl1, rl2)

}
