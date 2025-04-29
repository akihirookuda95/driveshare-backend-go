package main

import (
	"driveshare-backend-go/learning/calculator"
	"fmt"
)

func main() {
	fmt.Println("Hello, world!")

	fmt.Println(calculator.Offset)    // 2, public variable
	fmt.Println(calculator.Sum(1, 2)) // 4, public function
	// fmt.Println(calculator.offset) // error: offset is not exported by calculator

	fmt.Println("Multiply: ", calculator.Multiply(1, 2)) // 3, public function
}
