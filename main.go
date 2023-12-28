package main

import (
	"fmt"
)

func sayHello(name string) {
	fmt.Printf("Hello %s\n", name)
}

func add(a int, b int) int {
	return a + b
}

// Define the Student struct
type Student struct {
	firstName string
	lastName  string
}

// Method with a receiver of type Student
// This method returns the full name of the student
func (s Student) FullName() string {
	return s.firstName + " " + s.lastName
}

// Define a struct type
type Rectangle struct {
	Length float64
	Width  float64
}

// Method with a receiver of type Rectangle
func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}

// ----------------------------------------------------------------
// Speaker interface
type Speaker interface {
	Speak() string
}

// Dog struct
type Dog struct {
	Name string
}

// Dog's implementation of the Speaker interface
func (d Dog) Speak() string {
	return "Woof!"
}

// Person struct
type Person struct {
	Name string
}

// Person's implementation of the Speaker interface
func (p Person) Speak() string {
	return "Hello!"
}

// function that accepts Speaker interface
func makeSound(s Speaker) {
	fmt.Println(s.Speak())
}

func main() {
	sayHello("Shin")

	number1 := 3
	number2 := 5
	sumNumber := add(number1, number2)
	fmt.Println(sumNumber)

	student := Student{
		firstName: "Shin",
		lastName:  "TH",
	}

	// Call the FullName method on the Student instance
	fullName := student.FullName()
	fmt.Println("Full Name of the student:", fullName)

	rect := Rectangle{Length: 10, Width: 5}

	// Call the Area method on Rectangle instance
	area := rect.Area()
	fmt.Println("Area of rectangle:", area)

	//#Interface
	dog := Dog{Name: "Buddy"}
	person := Person{Name: "Alice"}

	makeSound(dog)
	makeSound(person)

}
