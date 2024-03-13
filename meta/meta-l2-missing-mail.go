package meta

func GetMaxExpectedProfit(N int32, V []int32, C int32, S float64) float64 {

	// start recurrence, calculate with and without picking email on day0

	// last day options:
	// - enter, pay C, get V(N) - C value + (1-S) * (whatever was before)
	// - don't enter
	// crtExpectedProfit := 0.0
	// profitPickMailLastDay := maxExpectedProfit(N, V, C, S, 0, crtProfit+float64(V[0]-C), 0)
	// profitLeaveMail := maxExpectedProfit(N, V, C, S, 0, crtProfit+(1-S)*float64(V[0]))

	// if profitPickMail > profitLeaveMail {
	// 	return profitPickMail
	// }

	return 0 // profitLeaveMail

	// evPrev := 0.0
	// collected := 0.0

	// for i := 0; i < int(N); i++ {
	// 	// just print the EVs
	// 	evCurrent := (1-S)*evPrev + float64(V[i]) - float64(C)*S
	// 	if (1-S)*evPrev > float64(C)*S && i > 0 {
	// 		// collect
	// 		fmt.Printf("Collect on day %d\n", i+1)
	// 		collected += evCurrent
	// 		// reset the days
	// 		evPrev = 0
	// 	} else {
	// 		// continue
	// 		evPrev = evCurrent
	// 	}
	// 	fmt.Printf("evCurrent: %f, (1-S)*ePrev: %f, float64(C)*S: %f\n", evCurrent, (1-S)*evPrev, float64(C)*S)
	// }

	// fmt.Printf("collected: %fn", collected)
	// return collected
}

// on any given day:
// - enter
//   - collected money == collect what's there (Vi - C) + what you collected before
//   - reset potential value to 0
//
// - don't enter
//   - collected money == only what you collected before
//   - add current value to potential value
// func maxExpectedProfit(N int32, V []int32, C int32, S float64, day int, currentProfit float64, expectedLeft float64) float64 {
// 	return 0
// }
