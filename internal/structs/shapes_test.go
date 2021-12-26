package main

import ("testing"
		"reflect"
)		

type Shape interface{
	Area() float64
}

func TestArea (t *testing.T) {
	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if !reflect.DeepEqual(got,want) {
			t.Errorf("got %g, want %g", got, want)
		}
	}
	t.Run("Rectangles", func(t *testing.T) {
		rectangle := Rectangle{40.0, 86.2}
		checkArea(t, rectangle, 3448)
	})
	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		checkArea(t, circle, 314.1592653589793)
	})
}