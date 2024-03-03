package meta

import (
	"fmt"
	"sort"
)

// Write any import statements here

// NOTE - "additional diners", not total (so not M + x)
// NOTE - doesn't matter how many at the end
// NOTE - currently social distancing is satisfied
// (S is not sorted) NOTE - is S sorted??
// 	- ... would help if sorted, consider that
// NOTE - careful with the seats starting at 1 (1 to N, same with S)
// (cheat) NOTE - int64 fits in an int on this machine
// NOTE - edge cases - return error?

var test = 1

// BASIC
func getMaxAdditionalDinersCount(N int64, K int64, M int32, S []int64) int64 {

	fmt.Printf("CASE %d: N = %d, K = %d, M = %d\n", test, N, K, M)
	fmt.Println(S)
	test++

	// some edge cases:
	// 	N, K, M, S - negatives ???
	if N <= 0 || K < 0 || N < int64(M) || int(M) != len(S) || len(S) > int(N) {
		return -1
	}
	// 	len S > N ???
	// 	N / M = K (can't put anyone anymore)
	//	...???

	// edge case - S is empty
	// . . a . . a . . a . . a .
	if M == 0 {
		return N / (K + 1)
	}

	// sort S in place
	sort.Slice(S, func(i, j int) bool { return S[i] < S[j] })
	sortedS := S
	fmt.Println(sortedS)

	// additional#
	additional := int64(0)

	// 1. before first, we need k + 1 multiples for #seats
	// 2. in between, a additionals would require a * (k + 1) + k for #seats
	// 3. after last, we need k + 1 multiples for #seats

	// 10, 1, 2, [2,6]
	// . x . . . x . . . .
	// 15, 2, 3, [6, 11, 14]
	// . . . . . x . . . . X . . X .

	// the math... parse sortedS
	// from 0 to sortedS[0]
	additional += (sortedS[0] - 1) / (K + 1)
	// fmt.Printf("possible at the beginning: %d\n", additional)
	// all from sortedS[i] to sortedS[i+1]
	for i := 0; i < int(M-1); i++ {
		possible := (sortedS[i+1] - sortedS[i] - 1 - K) / (K + 1)
		// fmt.Printf("possible in the middle: %d\n", possible)
		additional += possible
	}
	// from sorted[M-1] to N
	possible := (N - sortedS[M-1]) / (K + 1)
	// fmt.Printf("possible at the end: %d\n", possible)
	additional += possible

	return int64(additional)
}
