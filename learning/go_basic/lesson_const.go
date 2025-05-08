package main

const Pi = 3.14 // Pi is a constant, declared outside a function
const (
	Username = "admin" // Username is a constant, declared outside a function
	Password = "pass"  // Password is a constant, declared outside a function
)

var big int = 9223372036854775807 // the maximum value for int64

// var bigVar int = 9223372036854775807 + 1 // error: constant 9223372036854775808 overflows int
const bigConst = 9223372036854775807 + 1 // no error, constant is not evaluated at runtime

/*
func go_basic() {
	fmt.Println(Username, Password) // Username and Password are constants
	// Pi = 3 // error: cannot assign to Pi
	fmt.Println(big)     // 9223372036854775807
	fmt.Println(big + 1) // -9223372036854775808, overflow
	// Go is not occurs error on overflow and wraps minimum value.
	// Int range is -9223372036854775808 to 9223372036854775807

	fmt.Println(bigConst - 1) // 9223372036854775807, no error because constant is not evaluated at runtime

}

*/
