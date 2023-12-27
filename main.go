package main

import (
	"fmt"
)

func line() {
	fmt.Println("-----------------------------------")
}

func main() {
	// Declare variable
	// var fullname string = "Chinnawich Amapai"
	// var age int = 22
	fullName := `Chinnawich Ampai`
	age := 22

	fmt.Printf("Hello : %s\n", fullName)
	fmt.Printf("age : %d\n", age)

	fullName = "Shin"
	fmt.Println(fullName)
	line()

	//🚀 Type of Variable
	var booleanVar bool = true
	var intVar int = 10
	var floatVar float64 = 3.14
	var stringVar string = "Hello Go"

	fmt.Println("Boolean:", booleanVar)
	fmt.Println("Integer:", intVar)
	fmt.Println("Float:", floatVar)
	fmt.Println("String:", stringVar)
	line()

	// Constant
	const color string = "Red"
	fmt.Println("Color:", color)
	//Operator
	// use like another language
	a := 10
	b := 3
	fmt.Println(a + b) // 13
	fmt.Println(a - b) // 7
	fmt.Println(a * b) // 30
	fmt.Println(a / b) // 3
	fmt.Println(a % b) // 1
	line()

}
