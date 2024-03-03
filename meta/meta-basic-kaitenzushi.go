package meta

// BASIC
func isDishInLastEaten(lastAsMap map[int32]bool, dish int32) bool {
	found := lastAsMap[dish]
	return found
}

func addDishEaten(cutoff int32, lastAsSlice []int32, lastAsMap map[int32]bool, dish int32) []int32 {
	lastAsSlice = append(lastAsSlice, dish)
	lastAsMap[dish] = true
	if len(lastAsSlice) > int(cutoff) {
		toRemove := lastAsSlice[0]
		lastAsSlice = lastAsSlice[1:]
		delete(lastAsMap, toRemove)
	}
	return lastAsSlice
}

func getMaximumEatenDishCount(N int32, D []int32, K int32) int32 {

	lastEatenAsSlice := []int32{}
	lastEatenAsMap := make(map[int32]bool)
	totalEaten := 0

	for i := 0; i < int(N); i++ {
		currentDish := D[i]
		if !isDishInLastEaten(lastEatenAsMap, currentDish) {
			totalEaten++
			lastEatenAsSlice = addDishEaten(K, lastEatenAsSlice, lastEatenAsMap, currentDish)
		}
	}

	return int32(totalEaten)
}
