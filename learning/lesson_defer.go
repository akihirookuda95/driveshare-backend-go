package main

import (
	"fmt"
)

func foo() {
	defer fmt.Println("world", "foo") // after foo() returns

	fmt.Println("hello", "foo")
}

//func main() {
//	/*
//		foo()
//
//
//		// defer is used to execute a function after the surrounding function returns
//		defer fmt.Println("world") // after main() returns
//
//		fmt.Println("hello")
//
//	*/
//
//	/*
//		// stacking defer　最初に登録したものが最後に実行される
//		fmt.Println("run")
//		defer fmt.Println(1)
//		defer fmt.Println(2)
//		defer fmt.Println(3)
//		fmt.Println("success")
//
//	*/
//
//	file, _ := os.Open("./leargning/lesson_array.go")
//	defer file.Close()
//	data := make([]byte, 100) // bytes array of 100 bytes
//	file.Read(data)
//	fmt.Println(string(data))
//
//}
