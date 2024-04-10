package leetcode

/*
	NOTES:
	- the bodies of water are disjunct intervals between the taller elevations
	- going across the elevations
	- if elevation <= prev_elevation, nothing
	- if elevation > prev_elevation, opportunity to find new interval
		- go backwards until
			- found something
*/

func Trap(height []int) int {

	// go through the list
	// every step, pick the height and go BACK UNTIL same height or taller
	// - only those can change

	return -1
}
