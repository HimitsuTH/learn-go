package main

import (
	"fmt"
)

func main() {

	fullName := `Chinnawich Ampai`
	age := 22

	fmt.Printf("Hello : %s\n", fullName)
	fmt.Printf("age : %d\n", age)

	fullName = "Shin"
	fmt.Println(fullName)

	//🚀 Type of Variable
	var booleanVar bool = true
	var intVar int = 10
	var floatVar float64 = 3.14
	var stringVar string = "Hello Go"

	fmt.Println("Boolean:", booleanVar)
	fmt.Println("Integer:", intVar)
	fmt.Println("Float:", floatVar)
	fmt.Println("String:", stringVar)
}
