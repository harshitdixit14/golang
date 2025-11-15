package main

import "fmt"

func main() {
	fmt.Println("structs in golang")
	// there is no inheritance in golang; No super or parent

	harshit := User{"harshit", "harshit@go.dev", true, 23}

	fmt.Println("harshit user is : ", harshit)

	fmt.Printf("harshit details are: %+v\n", harshit)

	fmt.Printf("Name is: %v\n", harshit.Name)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
