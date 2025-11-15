package main

import "fmt"

func main() {
	fmt.Println("welcome to functions in golang")
	greeter()

	// not allowed to right function inside function in golang
	greeterTwo()

	result := adder(3, 5)

	fmt.Println("result is : ", result)

	proRes, message := proAdder(2, 3, 5, 6)
	fmt.Println("message is : ", message)
	fmt.Println("pro result is: ", proRes)
}

func adder(a int, b int) int { // mention return value as well
	return a + b
}

func proAdder(values ...int) (int, string) {
	total := 0
	for _, value := range values {
		total += value
	}
	return total, "hello"
}

func greeterTwo() {
	fmt.Println("anohter method")
}

func greeter() {
	fmt.Println("Namastey from golang")
}
