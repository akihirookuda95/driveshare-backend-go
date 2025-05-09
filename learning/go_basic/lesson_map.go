package main

//func go_basic() {
//	m := map[string]int{"apple": 1, "banana": 2, "orange": 3} // map of string to int, m is not nil so memory is allocated, then we can assign values
//	fmt.Println(m)                                            // map[apple:1 banana:2 orange:3]
//	fmt.Println(m["apple"])                                   // 1
//	m["banana"] = 300
//	fmt.Println(m) // map[apple:1 banana:300 orange:3]
//	m["new"] = 500
//	fmt.Println(m) // map[apple:1 banana:300 new:500 orange:3]
//
//	fmt.Println(m["nothing"]) // 0
//	v, ok := m["apple"]
//	fmt.Println(v, ok) // 1 true
//	v2, ok2 := m["nothing"]
//	fmt.Println(v2, ok2) // 0 false
//
//	m2 := make(map[string]int) // create a map of string to int, m2 is not nil so memory is allocated, then we can assign values
//	m2["pc"] = 5000
//	fmt.Println(m2) // map[pc:5000]
//
//	// var m3 map[string]int // declare a map of string to int, m3 is nil, so memory is not allocated, then we can not assign values
//	// m3["pc"] = 5000       // error: assignment to entry in nil map
//
//	var s []int // declare a slice of integers, s is nil, so memory is not allocated, then we can not assign values
//	if s == nil {
//		fmt.Println("s is nil") // this will be printed
//	} else {
//		fmt.Println("s is not nil")
//	}
//}
