package leetcode

func ProductExceptSelf(nums []int) []int {

	// 2 arrays
	// 1'st keeps the prefixes "forward products" - skips last
	// 2'nd keeps the suffixes "reverse products" - skips first

	var prefixes []int = make([]int, len(nums))
	var suffixes []int = make([]int, len(nums))
	var prods []int = make([]int, len(nums))
	var lastIndex = len(nums) - 1

	prefixes[0] = nums[0]
	for i := 1; i <= lastIndex; i++ {
		prefixes[i] = prefixes[i-1] * nums[i]
	}

	suffixes[lastIndex] = nums[lastIndex]
	for i := lastIndex - 1; i >= 0; i-- {
		suffixes[i] = suffixes[i+1] * nums[i]
	}

	for i := range len(nums) {
		if i == 0 { // first
			prods[i] = suffixes[i+1]
			continue
		}
		if i == lastIndex { // last
			prods[i] = prefixes[i-1]
			continue
		}

		prods[i] = prefixes[i-1] * suffixes[i+1]
	}

	return prods
}
