package main

import "math"

type Rectangle struct {
	width,
	height float64
}

func (r Rectangle) Area() float64{
	return r.width * r.height
}

type Circle struct {
	radius float64
}


func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

