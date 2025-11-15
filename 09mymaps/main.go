package main

import "fmt"

func main() {
	fmt.Println("welcome to maps class")

	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("List of all languages: ", languages)
	fmt.Println("JS shorts for : ", languages["JS"])

	delete(languages, "RB")

	fmt.Println("List of all languages: ", languages)

	// loops through maps

	for _, value := range languages { // _, value is comma ok syntax
		fmt.Printf("for key , value is %v\n", value)
	}
}
