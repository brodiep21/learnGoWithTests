package iteration

func Repeat(character string, numRun int) string {
	var repeated string
	for i := 0; i < numRun; i++ {
		repeated += character
	}
	return repeated
}
