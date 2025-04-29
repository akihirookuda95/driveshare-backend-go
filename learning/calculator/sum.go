package calculator

import "fmt"

var offset float64 = 1 // private variable, not exported, but accessible in the same package
var Offset float64 = 2 // public variable, exported, accessible in other packages

// Sum public function, exported, accessible in other packages
func Sum(a float64, b float64) float64 {
	fmt.Println("multiply: ", multiply(a, b))
	fmt.Println("Multiply: ", Multiply(a, b))
	return a + b + offset
}
