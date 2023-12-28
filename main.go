package main

import (
	"fmt"
)

func main() {
	x := 20
	pointerEx()

	fmt.Println(x) // Output: 50 (x is changed)
	changeValue(&x)

	employee()
}

func pointerEx() {
	// Declare an integer variable
	x := 10

	// Declare a pointer to an integer and assign it the address of x
	var p *int = &x

	// Print the value of x and the value at the pointer p
	fmt.Println("Value of x:", x)  // Output: Value of x: 10
	fmt.Println("Value at p:", *p) // Output: Value at p: 10

	// Modify the value at the pointer p
	*p = 20

	// x is modified since p points to x
	fmt.Println("New value of x:", x) // Output: New value of x: 20
}

func changeValue(ptr *int) {
	*ptr = 50
}

type Employee struct {
	Name   string
	Salary int
}

// Function to give a raise to an employee
func giveRaise(e *Employee, raise int) {
	e.Salary += raise
}

func employee() {
	emp := Employee{Name: "John Doe", Salary: 50000}

	giveRaise(&emp, 5000)
	fmt.Println("After raise:", emp)
}
