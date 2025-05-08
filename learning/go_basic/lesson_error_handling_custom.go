package main

import (
	"errors"
	// "fmt"
)

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("can't divide by zero")
	}
	return a / b, nil
}

//func go_basic() {
//	result, err := divide(10, 0)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	fmt.Println(result)
//}
