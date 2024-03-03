package meta

import "fmt"

func main_hit_probability() {

	fmt.Println(getWrongAnswers(0, "ABA"))

}

func getHitProbability(R int32, C int32, G [][]int32) float64 {

	possibleHits := 0

	for i := 0; i < int(R); i++ {
		for j := 0; j < int(C); j++ {
			if G[i][j] == 1 {
				possibleHits++
			}
		}
	}

	// Write your code here
	return float64(possibleHits) / float64(R*C)
}
