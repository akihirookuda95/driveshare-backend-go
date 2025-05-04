package main

//type Vertex struct {
//	X int // 大文字
//	Y int
//	// X Y int // short declaration
//	S string
//}

func changeVertex(v Vertex) {
	v.X = 1000
}

func changeVertex2(v *Vertex) {
	v.X = 1000 // equivalent to (*v).X = 1000, but automatically dereference
}

//func main() {
//	/*
//		v := Vertex{X: 1, Y: 2} // struct literal
//		fmt.Println(v)          // {1 2 }
//		fmt.Println(v.X, v.Y)   // 1 2
//		v.X = 100
//		fmt.Println(v.X, v.Y) // 100 2
//
//		v2 := Vertex{X: 1} // struct literal
//		fmt.Println(v2)    // {1 0 }, default value of Y is 0
//
//		v3 := Vertex{1, 2, "test"} // struct literal
//		fmt.Println(v3)            // {1 2 test}
//
//		v4 := Vertex{}                // struct literal
//		fmt.Println(v4)               // {0 0 } // default value of X, Y, S is 0, "", nil
//		fmt.Printf("%T %v\n", v4, v4) // {0 0 } // receive a struct
//
//		var v5 Vertex                 // struct literal
//		fmt.Println(v5)               // {0 0 } // default value of X, Y, S is 0, "", nil, not nil
//		fmt.Printf("%T %v\n", v5, v5) // {0 0 } // receive a struct
//
//		v6 := new(Vertex)             // struct literal
//		fmt.Println(v6)               // &{0 0 }
//		fmt.Printf("%T %v\n", v6, v6) // receive a pointer to struct
//
//		v7 := &Vertex{}               // struct literal
//		fmt.Println(v7)               // &{0 0 }
//		fmt.Printf("%T %v\n", v7, v7) // receive a pointer to struct
//
//		s := make([]int, 0)
//		s2 := []int{}
//		s3 := []int{1, 2, 3}
//		fmt.Println(s, s2, s3) // [] [] [1 2 3]
//
//	*/
//
//	v := Vertex{1, 2, "test"} // struct literal
//	changeVertex(v)
//	fmt.Println(v) // {1 2 test} // v is not changed, because v is passed by value
//
//	v2 := &Vertex{1, 2, "test"} // struct literal
//	changeVertex2(v2)
//	fmt.Println(v2)  // &{1000 2 test} // v2 is changed, because v2 is passed by reference
//	fmt.Println(*v2) // {1000 2 test} // dereference v2 to get the value of the struct
//
//}
