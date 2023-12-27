package main

import (
	"fmt"
)

func main() {
	a := 20
	b := 30

	// if eles
	if a == b {
		fmt.Println(10)
	} else {
		fmt.Println("Number must match.")
	}

	DayOfWeek()

	num1 := 5
	num2 := 10

	if sumNum := num1 + num2; sumNum >= 10 {
		fmt.Println("sumNum more than 10")
	}
}

func DayOfWeek() {

	// SWitch case
	var dayOfWeek = 3

	switch dayOfWeek {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Invalid Day")
	}
}
