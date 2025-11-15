package main

import (
	"fmt"
	"io"
	"net/http"
)

const url string = "https://cdis.iitk.ac.in/crs_app"

func main() {
	fmt.Println("web requests")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Println("response we got from url is : ", response)

	defer response.Body.Close() // caller's responsibility to close the connection

	databytes, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	content := string(databytes)

	fmt.Println(content)
}
