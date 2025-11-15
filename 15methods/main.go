package main

import "fmt"

func main() {
	fmt.Println("structs in golang")
	// there is no inheritance in golang; No super or parent

	harshit := User{"harshit", "harshit@go.dev", true, 23}

	fmt.Println("harshit user is : ", harshit)

	fmt.Printf("harshit details are: %+v\n", harshit)

	fmt.Printf("Name is: %v\n", harshit.Name)
	harshit.GetStatus()
	harshit.NewMail("test.dev@go.in")
	fmt.Printf("harshit details are: %+v\n", harshit)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
}

func (u User) NewMail(s string) {
	u.Email = s
	fmt.Println("New email is : ", u.Email) // mail not changed !!
}
