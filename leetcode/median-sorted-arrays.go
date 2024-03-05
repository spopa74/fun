package leetcode

// IN PROGRESS

/*
Situations:
- both arrays empty -> problem :)
- one empty -> find median of the other
- arrays are disjunct, a whole one is smaller than the other
	-> eliminate the smaller one #elements from the other
	-> then find median of the leftover
- arrays do "intersect" - have a common range == interesting
	-> 1 3 6 8 9
	->    5  8   10 12 14 17 20 22
	1. median first == 6, middle == 3'd
	2. median second == 13, middle == between 4'th and 5'th
	3. eliminate the short one, on both sides
	->     6 8 9
	->    5  8   10 12 14 17
	4. median first == 8, middle == 2'd
	5. median second == 11, middle == between 3'd and 4'th
	6. eliminate the short one, on both sides
	->     6 8 9 -> 1 on each side
	->       8   10 12 14 -> one on each side
	...again
	->       8 9 -> 1 on each side
	->       8   10 12 -> one on each side
	...again
	->         9 -> 1 on each side
	->       8   10 -> one on each side
	...again
	->         9 -> 1 on each side
	->          -> empty
	=> 9
	... OR ELIMINATE MAX(med1, med2) at every step

*/

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	// if both empty return crap
	if len(nums1) == 0 && len(nums2) == 0 {
		panic("Both arrays empty")
	}

	// if any empty return median of the other
	if len(nums1) == 0 {
		return Median(nums2, 0, len(nums2)-1)
	}
	if len(nums2) == 0 {
		return Median(nums1, 0, len(nums1)-1)
	}

	// if one completely smaller than another, eliminate from both sides
	// return median of leftover
	if nums1[len(nums1)-1] <= nums2[0] {
		return 1
	}

	for {
		return 1
	}

}

func Median(nums []int, startIndex int, endIndex int) float64 {
	l := endIndex - startIndex + 1
	if l%2 == 1 {
		return float64(nums[l/2+startIndex])
	}

	return float64(nums[l/2-1+startIndex]+nums[l/2+startIndex]) / 2
}
