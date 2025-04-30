package main

import "fmt"

func main() {
	b := []byte{72, 73}    // 72 and 73 are ascii code for H and I
	fmt.Println(b)         // [72 73]
	fmt.Println(string(b)) // HI

	c := []byte{'H', 'I'}  // 'H' and 'I' are characters
	fmt.Println(c)         // [72 73]
	fmt.Println(string(c)) // HI
}
