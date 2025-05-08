package main

import (
	"fmt"
	"os"
)

func readFile(filename string) error {
	_, err := os.Open(filename)
	if err != nil {
		// エラーに追加情報を付加して返す
		return fmt.Errorf("failed to open file %s: %w", filename, err)
	}
	return nil
}

//func main() {
//	err := readFile("non_existent_file.txt")
//	if err != nil {
//		fmt.Println(err)
//	}
//}
