package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("welcome to time study of golang")

	presentTime := time.Now()

	fmt.Println("present time is : ", presentTime)

	fmt.Println(presentTime.Format("01-02-2006 Monday, 15:04:05")) // mm-dd-yyyy

	createdDate := time.Date(2025, time.February, 19, 14, 45, 0, 0, time.UTC)

	fmt.Println("create date is : ", createdDate)
	fmt.Println(createdDate.Format("01-02-2006 Monday"))
}
