package main

// "go doc" command is used to view documentation
// example: go doc fmt Println
import (
	"fmt"
	"os/user" // Importing the user package using / for absolute path
	"time"
)

func main() {
	// Println is a public function in the fmt package
	fmt.Println("Hello, World!", time.Now()) // Now() is a public function in the time package
	fmt.Println(user.Current())              // Current() is a public function in the user package
}
