package main

import "fmt"

func main() {
	fmt.Println("welcome to array in golang")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Banana"
	fruitList[3] = "Carrot"

	fmt.Println("fruit list is : ", fruitList)
	fmt.Println("fruit list length is : ", len(fruitList))

	var vegList = [3]string{"potato", "beans", "mushroom"}

	fmt.Println("vegy list is : ", len(vegList))

}
