package main

func incrementGenerator() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

func circleArea(pi float64) func(radius float64) float64 {
	return func(radius float64) float64 {
		return pi * radius * radius
	}
}

//func go_basic() {
//	counter := incrementGenerator()
//	fmt.Println(counter()) // 1
//	fmt.Println(counter()) // 2
//	fmt.Println(counter()) // 3
//	fmt.Println(counter()) // 4
//
//	c1 := circleArea(3.14) // using closure when changing pi
//	fmt.Println(c1(2))     // 12.56
//
//	c2 := circleArea(3)
//	fmt.Println(c2(2)) // 12
//}
