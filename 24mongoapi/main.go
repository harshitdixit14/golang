package main

import (
	"fmt"
	"net/http"

	"github.com/harshitdixit14/mongoapi/router"
)

func main() {
	fmt.Println("mongodb api")
	r := router.Router()
	fmt.Println("Server is getting started...")
	http.ListenAndServe(":4000", r)
	fmt.Println("Listening at port 4000...")
}
