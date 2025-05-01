package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

// loggingSettings sets the logging settings for the program
func loggingSettings(logFile string) {
	logfile, _ := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	multiLogFile := io.MultiWriter(os.Stdout, logfile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(multiLogFile)
}

func main() {
	loggingSettings("learn.log")    // log file name
	_, err := os.Open("ldjlsfjlaj") // not exist file
	if err != nil {
		log.Println("error:", err) // log is a package that provides simple logging functions
		log.Fatalln("error:", err) // exits the program with a non-zero status code
	}

	log.Println("logging!")             // log is a package that provides simple logging functions
	log.Printf("%T %v", "test", "test") // string test, %T is the type of the variable, %v is the value of the variable

	log.Fatalf("%T %v", "test", "test") // string test
	log.Fatalln("error!")               // exits the program with a non-zero status code
	fmt.Println("OK!")                  // this line will not be executed because log.Fatalln() exits the program
}
