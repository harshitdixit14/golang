package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("enter a number: ")

	// comma ok or error ok syntax

	input, _ := reader.ReadString('\n')
	// input is try and _ is catch
	fmt.Println("your number is, ", input)
	fmt.Printf("your number is %T ", input)
}
