package leetcode

import "fmt"

func Candy(ratings []int) int {

	// stupid edge case - 1 kid
	if len(ratings) == 1 {
		return 1
	}

	// to help
	var lastIndex int = len(ratings) - 1
	fmt.Println("lastIndex: ", lastIndex)

	// first find max rating - then will allocate an array of counts (of ratings)
	var maxRating int = 0
	for i := 0; i <= lastIndex; i++ {
		if ratings[i] > maxRating {
			maxRating = ratings[i]
		}
	}
	fmt.Println("maxRating: ", maxRating)

	// counting sort type of processing
	// - go once through whole array
	// - find ratings, how many, where
	var countsRatings []int = make([]int, maxRating+1)
	var ratingsPositions map[int][]int = make(map[int][]int)
	for i := 0; i <= lastIndex; i++ {
		// i == position, ratings[i] == rating
		crtRating := ratings[i]
		countsRatings[crtRating]++
		positions, found := ratingsPositions[crtRating]
		if !found {
			ratingsPositions[crtRating] = []int{i}
		} else {
			positions = append(positions, i)
			ratingsPositions[crtRating] = positions
		}
	}

	fmt.Println("countsRatings: ", countsRatings)
	fmt.Println("ratingPositions: ", ratingsPositions)

	// at this point we know the ratings (counts) and where they are (ratingPositions)
	// parse counts, will hold ratings increasingly, "how many" with 0 rating, 1 rating, etc
	// the minim rating gets 1
	// the next rating check if it has neighbors, pick depending on neighbors
	// this will happen once per kid

	// candies list collects the #candy for each kid
	var candies []int = make([]int, len(ratings))
	for r := 0; r < len(countsRatings); r++ {
		if countsRatings[r] > 0 { // it means there are kids with this particular rating
			// pick the list of positions of kids with this rating
			listKids := ratingsPositions[r] // this should not be empty
			for i := 0; i < len(listKids); i++ {
				if listKids[i] == 0 { // if first, compare with second
					if candies[1] == 0 { // neighbor has 0
						candies[0] = 1
						continue
					}
					// if same rating, give same candies
					if ratings[0] == ratings[1] {
						candies[0] = 1 // apparently is ok to give less for same rating
						continue
					}
					// if not (rating is greater), give candies neighbor + 1
					candies[0] = candies[1] + 1

				} else if listKids[i] == lastIndex { // last, compare with prev
					// similar with first
					if candies[lastIndex-1] == 0 {
						candies[lastIndex] = 1
						continue
					}
					// if same rating, give same candies
					if ratings[lastIndex] == ratings[lastIndex-1] {
						candies[lastIndex] = 1 // apparently is ok to give less for same rating
						continue
					}
					// if not (rating is greater), give candies neighbor + 1
					candies[lastIndex] = candies[lastIndex-1] + 1

				} else {
					// if neither... somewhere in the middle
					crtPosition := listKids[i]
					// 0 ? 0
					if candies[crtPosition-1] == 0 && candies[crtPosition+1] == 0 {
						candies[crtPosition] = 1
						continue
					}
					// 0 ? r
					if candies[crtPosition-1] == 0 && candies[crtPosition+1] != 0 {
						if ratings[crtPosition] == ratings[crtPosition+1] {
							candies[crtPosition] = 1 // apparently is ok to give less for same rating
							continue
						}
						// rating is greater
						candies[crtPosition] = candies[crtPosition+1] + 1
						continue
					}
					// l ? 0
					if candies[crtPosition-1] != 0 && candies[crtPosition+1] == 0 {
						if ratings[crtPosition] == ratings[crtPosition-1] {
							candies[crtPosition] = 1 // apparently is ok to give less for same rating
							continue
						}
						// rating is greater
						candies[crtPosition] = candies[crtPosition-1] + 1
						continue
					}
					// none of neighbors is 0
					// l = ? = r
					if ratings[crtPosition-1] == ratings[crtPosition] && ratings[crtPosition] == ratings[crtPosition+1] {
						candies[crtPosition] = 1 // apparently allowed
						continue
					}
					// l = ? < r
					if ratings[crtPosition-1] == ratings[crtPosition] && ratings[crtPosition] < ratings[crtPosition+1] {
						candies[crtPosition] = 1 // apparently allowed
						continue
					}
					// l = ? > r
					if ratings[crtPosition-1] == ratings[crtPosition] && ratings[crtPosition] > ratings[crtPosition+1] {
						candies[crtPosition] = candies[crtPosition+1] + 1
						continue
					}

					// l < ? = r
					if ratings[crtPosition-1] < ratings[crtPosition] && ratings[crtPosition] == ratings[crtPosition+1] {
						candies[crtPosition] = candies[crtPosition-1] + 1
						continue
					}
					// l < ? < r
					if ratings[crtPosition-1] < ratings[crtPosition] && ratings[crtPosition] < ratings[crtPosition+1] {
						candies[crtPosition] = candies[crtPosition-1] + 1
						continue
					}
					// l < ? > r
					if ratings[crtPosition-1] < ratings[crtPosition] && ratings[crtPosition] > ratings[crtPosition+1] {
						if candies[crtPosition-1] < candies[crtPosition+1] {
							candies[crtPosition] = candies[crtPosition+1] + 1
							continue
						}
						candies[crtPosition] = candies[crtPosition-1] + 1
						continue
					}
					// l > ? = r
					if ratings[crtPosition-1] > ratings[crtPosition] && ratings[crtPosition] == ratings[crtPosition+1] {
						candies[crtPosition] = 1
						continue
					}
					// l > ? < r
					if ratings[crtPosition-1] > ratings[crtPosition] && ratings[crtPosition] < ratings[crtPosition+1] {
						candies[crtPosition] = 1
						continue
					}
					// l > ? > r
					if ratings[crtPosition-1] > ratings[crtPosition] && ratings[crtPosition] > ratings[crtPosition+1] {
						candies[crtPosition] = candies[crtPosition+1] + 1
						continue
					}
				}
			}
		}
	}
	// debug
	fmt.Println("candies: ", candies)

	// find sum of candies
	sum := 0
	for i := 0; i <= lastIndex; i++ {
		sum += candies[i]
	}

	return sum
}
