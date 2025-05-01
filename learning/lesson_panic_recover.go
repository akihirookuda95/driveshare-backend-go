package main

import "fmt"

func thirdPartyConnectDB() {
	panic("Unable to connect to third party database") // panic is used to indicate that something went wrong
	// panicを使うことは推奨されてない
}

func save() {
	thirdPartyConnectDB()
	defer func() {
		s := recover()
		fmt.Println("recover", s) // recover is used to catch the panic and continue the program.
	}()
}

func main() {
	save()

	fmt.Println("ok")

}
