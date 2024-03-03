package leetcode

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// if both empty return crap
	if len(nums1) == 0 && len(nums2) == 0 {
		panic("Both arrays empty")
	}

	if len(nums1) == 0 {
		return Median(nums2)
	}

	if len(nums2) == 0 {
		return Median(nums1)
	}

	return 0
}

func Median(nums []int) float64 {
	l := len(nums)
	if l%2 == 1 {
		return float64(nums[l/2])
	}

	return float64(nums[l/2-1]+nums[l/2]) / 2
}
