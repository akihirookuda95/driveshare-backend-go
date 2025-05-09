package main

func getOsName() string {
	return "mac"
}

//func go_basic() {
//	// normal switch statement, os variable is usable outside and inside the switch statement
//	//os := getOsName()
//	//switch os {
//	//case "linux":
//	//	fmt.Println("Linux")
//	//case "windows":
//	//	fmt.Println("Windows")
//	//case "mac":
//	//	fmt.Println("Mac")
//	//default:
//	//	fmt.Println("Unknown OS")
//	//}
//
//	// simple switch statement, os variable is usable only inside the switch statement
//	switch os := getOsName(); os {
//	case "linux":
//		fmt.Println("Linux")
//	case "windows":
//		fmt.Println("Windows")
//	case "mac":
//		fmt.Println("Mac")
//	default:
//		fmt.Println("Unknown OS, as is", os)
//	}
//
//	t := time.Now()
//	fmt.Println(t.Hour())
//	switch {
//	case t.Hour() < 12:
//		fmt.Println("Good morning")
//	case t.Hour() < 18:
//		fmt.Println("Good afternoon")
//	}
//
//}
