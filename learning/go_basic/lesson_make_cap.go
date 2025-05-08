package main

//func go_basic() {
//	n := make([]int, 3, 5)                                    // create a slice of 3 integers with a capacity of 5
//	fmt.Println(n)                                            // [0 0 0]
//	fmt.Printf("len=%d cap=%d value=%v\n", len(n), cap(n), n) // len=3 cap=5 value=[0 0 0]
//	n = append(n, 0, 0)                                       // append 2 more elements, to append multiple elements, use n = append(n, 0, 0)
//	fmt.Printf("len=%d cap=%d value=%v\n", len(n), cap(n), n) // len=5 cap=5 value=[0 0 0 0 0]
//	n = append(n, 0, 0, 0, 0, 0)                              // append 5 more elements
//	fmt.Printf("len=%d cap=%d value=%v\n", len(n), cap(n), n) // len=10 cap=10 value=[0 0 0 0 0 0 0 0 0 0]
//
//	a := make([]int, 3)                                       // create a slice of 3 integers with a capacity of 3
//	fmt.Printf("len=%d cap=%d value=%v\n", len(a), cap(a), a) // len=3 cap=3 value=[0 0 0]
//
//	b := make([]int, 0)                                       // create a slice of 0 integers with a capacity of 0, b is not nil because it is initialized with make
//	var c []int                                               // create a slice of 0 integers with a capacity of 0, c is nil because it is not initialized with make
//	fmt.Printf("len=%d cap=%d value=%v\n", len(b), cap(b), b) // len=3 cap=3 value=[]
//	fmt.Printf("len=%d cap=%d value=%v\n", len(c), cap(c), c) // len=0 cap=0 value=[]
//
//	// c = make([]int, 3) // create a slice of 3 integers with a capacity of 3
//	c = make([]int, 0, 3) // create a slice of 0 integers with a capacity of 3
//	for i := 0; i < len(c); i++ {
//		c = append(c, i)
//		fmt.Println(c)
//	}
//	fmt.Println(c)
//}
