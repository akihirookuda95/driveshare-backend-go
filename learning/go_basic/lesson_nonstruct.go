package main

// original int type MyInt
type MyInt int

func (i MyInt) Double() int {
	return int(i * 2)
}

//func go_basic() {
//	myInt := MyInt(10)
//	fmt.Println(myInt.Double())
//
//}
