package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to web verb video")

	// PerformGetRequest()

	// PerformPostJsonRequest()

	PerformPostFormRequest()

}

func PerformGetRequest() {
	const myurl = "http://localhost:3000/get"

	response, err := http.Get(myurl)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("response is : ", response.StatusCode)
	fmt.Println("content length is : ", response.ContentLength)

	var responseString strings.Builder
	content, _ := io.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)

	fmt.Println("Byte count is : ", byteCount)
	fmt.Println(responseString.String())

	// fmt.Println(string(content))

	// same way to convert content to string initially content is in bytes

}

func PerformPostJsonRequest() {
	const myurl = "http://localhost:3000/post"

	// fake json payload

	requestBody := strings.NewReader(`
		{
			"coursename":"let's go with golang",
			"price":0,
			"platform":"udemy"
		}
	`)

	response, err := http.Post(myurl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, _ := io.ReadAll(response.Body)
	fmt.Println(string(content))
}

func PerformPostFormRequest() {
	const myurl = "http://localhost:3000/postform"

	// create form data

	data := url.Values{}

	data.Add("firstname", "Harshit")
	data.Add("lastname", "Dixit")
	data.Add("email", "harshit@go.dev")

	response, err := http.PostForm(myurl, data)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, _ := io.ReadAll(response.Body)

	fmt.Println(string(content))

}
