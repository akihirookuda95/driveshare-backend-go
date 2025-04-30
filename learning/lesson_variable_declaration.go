package main

var a int = 1 // a is a global variable, declared outside a function
// b := 2 // error: b is not declared outside a function

//func foo() {
//	// short variable declaration using inside only a function
//	xi := 1         // xi is a shorthand for var xi int = 1
//	fmt.Println(xi) // 1
//	xi = 2
//	fmt.Println(xi) // 2
//	var xf32 float32 = 1.2
//	xs := "Hello, World!"             // xs is a shorthand for var xs string = "Hello, World!"
//	xt, xf := true, false             // xt is a shorthand for var xt, xf bool = true, false
//	fmt.Println(xi, xf32, xs, xt, xf) // 1 1.2 Hello, World! true false
//	fmt.Printf("%T\n", xf32)          // float32, %T is a format specifier for type
//	fmt.Printf("%T\n", xi)            // int, \n is a format specifier for new line
//}

/*
func main() {
	// parenthesis () are used to group expressions inside and outside a function
	var (
		i    int     = 1
		f64  float64 = 1.2
		s    string  = "Hello, World!"
		t, f bool    = true, false // using , to declare multiple variables
	)
	fmt.Println(i, f64, s, t, f)

	//var (
	//	i    int     // 0 is the default value for int
	//	f64  float64 // 0.0 is the default value for float64
	//	s    string  // "" is the default value for string
	//	t, f bool    // false is the default value for bool
	//)
	//fmt.Println(i, f64, s, t, f)

	// short variable declaration using inside only a function
	xi := 1               // xi is a shorthand for var xi int = 1
	xf64 := 1.2           // xf64 is a shorthand for var xf64 float64 = 1.2
	xs := "Hello, World!" // xs is a shorthand for var xs string = "Hello, World!"
	xt, xf := true, false // xt is a shorthand for var xt, xf bool = true, false
	fmt.Println(xi, xf64, xs, xt, xf)
	fmt.Printf("%T\n", xf64) // float64, %T is a format specifier for type

	//fmt.Println(a) // 1, a is a global variable

	foo()
}
*/
