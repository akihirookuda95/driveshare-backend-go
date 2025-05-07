package main

// Receiver is mechanism to attach a method to a type.
// It can be a value receiver or a pointer receiver.
// Value receiver: a copy of the value is passed to the method.
// Pointer receiver: a pointer to the value is passed to the method.

type ValueExample struct {
	Value int
}

// This receiver is a value receiver,
// which means that a copy of the ValueExample struct is passed to the method.
func (v ValueExample) Increment() {
	v.Value += 1
}

type PointerExample struct {
	Value int
}

// This receiver is a pointer receiver,
// which means that a pointer to the PointerExample struct is passed to the method.
func (p *PointerExample) Increment() {
	p.Value += 1
}

//func main() {
//	v := ValueExample{Value: 42}
//	v.Increment()  // This will not change the original value of v
//	fmt.Println(v) // Output: {42}
//
//	p := PointerExample{Value: 42}
//	p.Increment()  // This will change the original value of p
//	fmt.Println(p) // Output: &{43}
//}
