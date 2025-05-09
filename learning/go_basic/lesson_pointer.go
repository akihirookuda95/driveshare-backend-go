package main

func one(x *int) {
	*x = 1
}

//func go_basic() {
//	var n int = 100 // declare a variable n of type int and assign it a value of 100
//	one(&n)         // dereference the pointer to n and assign it a value of 1
//	fmt.Println(n)  // 1
//
//	//fmt.Println(n)  // 100
//	//fmt.Println(&n) // address of n
//	//
//	//var p *int = &n // declare a pointer variable p of type *int and assign it the address of n
//	//fmt.Println(p)  // address of n
//	//fmt.Println(*p) // 100
//}
