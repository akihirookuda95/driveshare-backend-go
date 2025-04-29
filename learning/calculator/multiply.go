package calculator

// Multiply public function, exported, accessible in other packages
func Multiply(a float64, b float64) float64 {
	return (a * b) + offset
}

// multiply private function, not exported, accessible only in the same package
func multiply(a float64, b float64) float64 {
	return (a * b) + offset
}
