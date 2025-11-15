package main

import "fmt"

func main() {
	fmt.Println("welcome to loops in golang")

	// days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// for i := range days { // i is index here !!
	// 	fmt.Println(days[i])
	// }

	// for index, day := range days {
	// 	fmt.Printf("index is %v and value is %v: \n", index, day)
	// }

	rougueValue := 1

	for rougueValue < 10 {
		if rougueValue == 2 {
			goto lco
		}
		if rougueValue == 5 {
			rougueValue++
			continue
		}
		fmt.Printf("value is %v\n", rougueValue)
		rougueValue++
	}

	// break and continue

lco:
	fmt.Println("jumping at lco")

}
