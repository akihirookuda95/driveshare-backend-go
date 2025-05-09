package main

//func go_basic() {
//	/*
//		var p *int = new(int) // declare a pointer variable p of type *int and assign it a new int
//		fmt.Println(p)        // address of the new int, allocated in heap memory
//		*p++                  // dereference the pointer to p and assign it a value of 100
//		fmt.Println(*p)       // 1, the value of the new int is 0
//
//		var p2 *int
//		fmt.Println(p2) // <nil>, p2 is nil, so memory is not allocated, then we can not assign values
//		*p2++
//		// fmt.Println(p2) // error: invalid memory address or nil pointer dereference
//
//	*/
//
//	// using make is to allocate memory for a slice, map, or channel
//	// using new is to allocate memory for a pointer or struct
//
//	s := make([]int, 0)
//	fmt.Printf("%T\n", s) // []int, slice of int
//
//	m := make(map[string]int)
//	fmt.Printf("%T\n", m) // map[string]int, map of string to int
//
//	ch := make(chan int)
//	fmt.Printf("%T\n", ch) // chan int, channel of int
//
//	var p *int = new(int)
//	fmt.Printf("%T\n", p) // *int, pointer to int
//
//	var st = new(struct{})
//	fmt.Printf("%T\n", st) // *struct {}, pointer to struct
//
//}
