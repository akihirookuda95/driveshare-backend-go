package main

import (
	"fmt"
)

type Vertex struct {
	X, Y int // coordinates
}

func (v Vertex) Area() int {
	return v.X * v.Y
}

// pointer receiver
func (v *Vertex) Scale(i int) {
	v.X = v.X * i
	v.Y = v.Y * i
}

// variable receiver
func Area(v Vertex) int {
	return v.X * v.Y
}

func main() {
	v := Vertex{3, 4}
	// fmt.Println(Area(v)) // function receiver

	v.Scale(10) // pointer receiver

	fmt.Println(v.Area()) // method receiver
}
