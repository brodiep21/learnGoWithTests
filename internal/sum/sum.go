package main

func Sum(numbers []int) int {
	sum := 0
	for _, i := range numbers {
		sum += i
	}
	return sum
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return sums
}
func SumHeads(variableSlices ...[]int) []int {
	var number []int
	for _, i := range variableSlices {
		if len(i) == 0 {
			number = append(number, 0)
		} else {
			head := i[:1]
			number = append(number, Sum(head))
		}
	}
	return number
}

// 	 //talk to eli or lookup issues with empty slice pulling out of bounds error - panic: runtime error: slice bounds out of range [:1] with capacity 0 [recovered]
// panic: runtime error: slice bounds out of range [:1] with capacity 0
