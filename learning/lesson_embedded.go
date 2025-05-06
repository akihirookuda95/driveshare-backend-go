package main

type Vertex struct {
	x, y int // coordinates
}

// variable receiver
// calling method, send the copy of the struct
// not modify the original struct
// variable receiver is used when the method does not need to modify the original struct
// and when copy cost is low
func (v Vertex) Area() int {
	return v.x * v.y
}

// pointer receiver
// calling method, send the pointer of the struct
// modiqy the original struct
// ponter receiver is used when the method needs to modify the original struct
// and do not want to cost of copy of the struct
func (v *Vertex) Scale(i int) {
	v.x = v.x * i
	v.y = v.y * i
}

type Vertex3D struct {
	Vertex
	z int
}

func (v Vertex3D) Area3D() int {
	return v.x * v.y * v.z
}

func (v *Vertex3D) Scale3D(i int) {
	v.x = v.x * i
	v.y = v.y * i
	v.z = v.z * i
}

// constructor
// "New" is a convention to name the constructor function
// Return a pointer to the struct, not a copy of the struct, so effectively
func New(x, y, z int) *Vertex3D {
	return &Vertex3D{Vertex{x, y}, z}
}

//func main() {
//	v := New(3, 4, 5)
//	v.Scale(10)
//	fmt.Println(v.Area())   // 1200
//	fmt.Println(v.Area3D()) // 6000
//}
