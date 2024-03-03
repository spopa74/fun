package main

import (
	"errors"
	"fmt"
	"unicode"
)

// WARMUP
func nextLine(s string) (string, error) {

	// if param empty, error
	if s == "" {
		return "", errors.New("param is an empty string")
	}

	sliceDigits := []rune{}
	sliceCounts := []int{}
	sliceIndex := -1
	lastChar := ' '

	for i := 0; i < len(s); i++ {
		c := rune(s[i])
		// if it runs into a non-digit, error
		if !unicode.IsDigit(c) {
			errMessage := fmt.Sprintf("ran into a non-digit character, char %c, position %d.", c, i)
			return "", errors.New(errMessage)
		}

		if c == lastChar {
			// increment sliceCount
			sliceCounts[sliceIndex]++
		} else {
			sliceIndex++
			sliceDigits = append(sliceDigits, c)
			sliceCounts = append(sliceCounts, 1)
			lastChar = c
		}
	}

	result := ""
	for i := 0; i < len(sliceDigits); i++ {
		result += fmt.Sprintf("%d", sliceCounts[i])
		result += string(sliceDigits[i])
	}

	return result, nil
}
