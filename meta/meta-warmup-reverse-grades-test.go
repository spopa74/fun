package meta

// WARMUP
// assumes indeed N >=0 and len(C) == N
func GetWrongAnswers(N int32, C string) string {

	var response string

	for i := 0; i < int(N); i++ {
		if C[i] == 'A' {
			response += "B"
		} else {
			response += "A"
		}
	}

	return response
}
