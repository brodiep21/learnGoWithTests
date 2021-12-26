package main

import ("testing"
		"reflect"
)

func TestSum (t *testing.T) {

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1,2,3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

}
func TestSumAllTails(t *testing.T) {

	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}
	t.Run("make the sums of tails of", func(t *testing.T){
		got := SumAllTails([]int{1,2}, []int{0,9}, []int{4,9,16})
		want := []int{2, 9, 25}
		checkSums(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{4,7,0,5})
		want := []int{0, 12}
		checkSums(t, got, want)
	})
}

func TestHead(t *testing.T) {
	got := SumHeads([]int{3,4,6,88,9}, []int{1,222,45})
	want := []int{3,1}
	
	if !reflect.DeepEqual(got,want) {
		t.Errorf("got %v, want %v", got, want)
	}
}