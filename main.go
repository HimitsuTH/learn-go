package main

import (
	"fmt"

	greetings "github.com/HimitsuTH/learn-go/sayHello"
	"github.com/google/uuid"
)

func main() {
	id := uuid.New()
	fmt.Println("Hello world 3")
	fmt.Printf("UUID: %s\n", id)
	greetings.SayHello()
	greetings.SayTest()
}
