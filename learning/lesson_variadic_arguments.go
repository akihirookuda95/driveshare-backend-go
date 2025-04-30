package main

import "fmt"

// function having variadic arguments, ... is used to indicate that the function can take a variable number of arguments
func foo(params ...int) {
	fmt.Println(len(params), params)
	for _, param := range params {
		fmt.Println(param)
	}
}

//func main() {
//	foo()           // 0 []
//	foo(10, 20)     // 2 [10 20] 10 20
//	foo(10, 20, 30) // 3 [10 20 30] 10 20 30
//
//	s := []int{1, 2, 3}
//	fmt.Println(s) // [1 2 3]
//	foo(s...)      // 3 [1 2 3] 1 2 3, ... is used to unpack the slice
//	// foo(s) // error: cannot use s (type []int) as type int in argument to foo
//}
