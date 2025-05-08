//package main
//
//import (
//	"errors"
//	"fmt"
//)
//
//// Custom error type
//var ErrNotFound = errors.New("not found")
//
//func findItem(id int) error {
//	if id != 1 {
//		return ErrNotFound
//	}
//	return nil
//}
//
//func main() {
//	err := findItem(1)
//	if errors.Is(err, ErrNotFound) {
//		fmt.Println("Item not found")
//	} else if err != nil {
//		fmt.Println("An error occurred:", err)
//	}
//	fmt.Println("Item found")
//}

package main

import (
	"errors"
	"fmt"
)

type MyError struct {
	Message string
}

func (e *MyError) Error() string {
	return e.Message
}

func doSomething() error {
	return &MyError{Message: "something went wrong"}
}

func main() {
	err := doSomething()
	var myErr *MyError
	if errors.As(err, &myErr) {
		fmt.Println("Custom error:", myErr.Message)
	} else if err != nil {
		fmt.Println("An error occurred:", err)
	}
}
