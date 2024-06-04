package leetcode

func ValidPalindrome(s string) bool {
	length := len(s)
	for i := 0; i <= length/2; i++ {
		if s[i] != s[length-1-i] {
			// check palindrome removing right
			left := checkPalindrome(s[i : length-1-i])
			// check palindrome removing left
			right := checkPalindrome(s[i+1 : length-i])

			if left || right {
				return true
			} else {
				return false
			}

		}
	}

	return true
}

func checkPalindrome(s string) bool {
	length := len(s)
	for i := 0; i <= length/2; i++ {
		if s[i] != s[length-1-i] {
			return false
		}
	}

	return true
}
