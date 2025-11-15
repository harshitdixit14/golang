package main

import "fmt"

const LoginToken string = "hello"

func main() {
	var username string = "hello"
	fmt.Println(username)
	fmt.Printf("varibale is of type : %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("varibale is of type : %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("varibale is of type : %T \n", smallVal)

	var smallFloat float64 = 255.453487342423
	fmt.Println(smallFloat)
	fmt.Printf("varibale is of type : %T \n", smallFloat)

	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("varibale is of type : %T \n", anotherVariable)

	var anotherVariable2 string
	fmt.Println(anotherVariable2)
	fmt.Printf("varibale is of type : %T \n", anotherVariable2)

	// implicit type

	var website = "learncodeonline.in"
	fmt.Println(website)

	numberOfUser := 300000
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
}
