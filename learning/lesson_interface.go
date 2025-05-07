package main

import (
	"fmt"
)

// interface
type Animal interface {
	Speak() string
	Rename(newName string)
	GetName() string
}

// Dog struct
type Dog struct {
	Name string
}

// variable receiver: will not change the original value
func (d Dog) Speak() string {
	return d.Name + " says Woof!"
}

// pointer receiver: will change the original value
func (d *Dog) Rename(newName string) {
	d.Name = newName
}

// variable receiver: will not change the original value
func (d Dog) GetName() string {
	return d.Name
}

// Cat struct
type Cat struct {
	Name string
}

// variable receiver: will not change the original value
func (c Cat) Speak() string {
	return c.Name + " says Meow!"
}

// pointer receiver: will change the original value
func (c *Cat) Rename(newName string) {
	c.Name = newName
}

// variable receiver: will not change the original value
func (c Cat) GetName() string {
	return c.Name
}

// common code
func Introduce(animal Animal) {
	fmt.Println(animal.Speak())
	animal.Rename("New-Name")
	fmt.Println("New name: ", animal.GetName())
}

func main() {
	dog := &Dog{Name: "Max"}
	cat := &Cat{Name: "Whiskers"}

	Introduce(dog)
	Introduce(cat)
}
