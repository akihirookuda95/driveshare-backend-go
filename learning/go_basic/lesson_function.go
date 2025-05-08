package main

func add(x int, y int) (int, int) {
	return x + y, x - y
}

func cal(price int, item int) (result int) {
	result = price * item // result is declared as argument of the function, so it is not necessary to declare it again
	return                // equivalent to return result
}

func convert(price int) float64 {
	return float64(price) // convert int to float64
}

//func go_basic() {
//	r, r2 := add(1, 2) // initialize r with the return value of add
//	fmt.Println(r)     // 3
//	fmt.Println(r2)    // -1
//	// fmt.Println(add("hello", 2)) // error: cannot use "hello" (type untyped string) as type int in argument to add
//
//	r3 := cal(100, 2)
//	fmt.Println(r3) // 200
//
//	r4 := convert(100)
//	fmt.Println(r4) // 100.000000
//
//	f := func(x int) {
//		fmt.Println("inner function", x)
//	} // anonymous function
//	f(1) // inner function 1
//
//	func(x int) {
//		fmt.Println("inner function", x)
//	}(2) // inner function 2
//
//}
