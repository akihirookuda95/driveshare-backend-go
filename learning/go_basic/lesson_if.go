package main

func by2(num int) string {
	if num%2 == 0 {
		return "ok"
	} else {
		return "not ok"
	}
}

//func go_basic() {
//
//	result := by2(10)
//	fmt.Println(result)
//	if result == "ok" {
//		fmt.Println("by 2")
//	}
//
//	// simplified if statement
//	if result2 := by2(10); result2 == "ok" {
//		fmt.Println("by 2")
//	}
//	// fmt.Println(result2) // error: result2 is not defined outside the if statement
//
//	num := 6
//	if num%2 == 0 {
//		fmt.Println("by 2")
//	} else if num%3 == 0 {
//		fmt.Println("by 3")
//	} else {
//		fmt.Println("else")
//	}
//
//	x, y := 10, 10
//	if x == 10 && y == 10 {
//		fmt.Println("x and y are both 10")
//	}
//
//	if x == 10 || y == 10 {
//		fmt.Println("x or y is 10")
//	}
//
//}
