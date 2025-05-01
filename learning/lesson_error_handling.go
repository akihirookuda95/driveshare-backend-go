package main

//func main() {
//	file, err := os.Open("./lesson_array.go") // Both file and err are initialized with the return value of os.Open
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer file.Close()
//	data := make([]byte, 100)
//	// short variable declaration, when two variables are inserted, if the other variable is initialized, we can update the other variable.
//	count, err := file.Read(data) // count variable is initialized with the number of bytes read, err variable is updated with the error
//	if err != nil {
//		log.Fatal("error:", err)
//	}
//	fmt.Println(count, string(data))
//
//	//err := os.Chdir("test") // error no new variable on left side
//	//if err != nil {
//	//	log.Fatal(err)
//	//}
//
//	// err variable is used inside the if statement, so we can use the same variable name as the one outside the if statement
//	if err = os.Chdir("./"); err != nil {
//		log.Fatalln(fmt.Errorf("could not change dir: %v", err))
//	}
//
//}
