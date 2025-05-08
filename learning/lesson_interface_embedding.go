package main

import (
	"fmt"
)

// 基本的なインターフェース
type Reader interface {
	Read() string
}

type Writer interface {
	Write(data string)
}

// ReaderとWriterを埋め込んだ新しいインターフェース
type ReadWriter interface {
	Reader
	Writer
}

// 実装
type File struct {
	Content string
}

func (f *File) Read() string {
	return f.Content
}

func (f *File) Write(data string) {
	f.Content = data
}

func main() {
	var file ReadWriter = &File{}

	file.Write("Hello, Go!")
	fmt.Println(file.Read()) // Hello, Go!

}
