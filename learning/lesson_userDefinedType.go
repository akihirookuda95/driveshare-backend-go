// User defined types are
// 1. new types based on existing types
// 2. struct types
// 3. new types based on slice, map and channel
// 4. interface
// 5. pointer types

package main

import (
	"fmt"
)

// 1. new types based on existing types
type UserID int

func GetUserNameByID(id UserID) string {
	return fmt.Sprintf("User ID: %d", id)
}

// 2. struct types
type User struct {
	ID   UserID
	Name string
	Age  int
}

func (u User) IsAdult() bool {
	return u.Age > 18
}

// 3. new types based on slice, map and channel
type ProductList []string

func (p ProductList) PrintAll() {
	for _, product := range p {
		fmt.Println(product)
	}
}

// 4. interface
type Animal interface {
	Speak() string
}

type Dog struct {
	Name string
}

func (d Dog) Speak() string {
	return d.Name + " says Woof!"
}

type Cat struct {
	Name string
}

func (c Cat) Speak() string {
	return c.Name + " says Meow!"
}

// 5. pointer types
type IntPointer *int

func Increment(p IntPointer) {
	*p++
}

func main() {
	// 1. new types based on existing types
	var id UserID = 101
	fmt.Println(GetUserNameByID(id)) // User ID: 101

	// 2. struct types
	user := User{ID: id, Name: "John Doe", Age: 25}
	fmt.Println(user.IsAdult()) // true

	// 3. new types based on slice, map and channel
	products := ProductList{"Laptop", "Smartphone", "Tablet"}
	products.PrintAll() // Laptop Smartphone Tablet

	// 4. interface
	animals := []Animal{
		Dog{Name: "Buddy"},
		Cat{Name: "Whiskers"},
	}
	for _, animal := range animals {
		fmt.Println(animal.Speak()) // Buddy says Woof! Whiskers says Meow!
	}

	// 5. pointer types
	var value int = 10
	var p IntPointer = &value
	Increment(p)
	fmt.Println(*p) // 11
}
