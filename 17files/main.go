package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("welcome to files in golang")
	// text files
	content := "This needs to go in a file - harshit dixit"

	file, err := os.Create("./myfile.txt")

	if err != nil {
		panic(err) // panic shuts the execution of the program and shows the error
	}
	// writing file
	length, err := io.WriteString(file, content) // read documentation that WriteString return length

	if err != nil {
		panic(err)
	}
	fmt.Println("length is : ", length)
	defer file.Close()
	// reading file
	readFile("myfile.txt")
}

func readFile(filename string) {
	databyte, err := os.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	fmt.Println("text data inside the file is : ", string(databyte))
}
