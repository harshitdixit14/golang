package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://leetcode.com/problems/transform-to-chessboard/description/?envType=problem-list-v2&envId=bit-manipulation"

func main() {
	fmt.Println("url handling in golang")
	fmt.Println(myurl)

	// parsing url

	// handling url

	result, _ := url.Parse(myurl)

	// fmt.Println(result.Scheme)
	// fmt.Println(result.Host)
	// fmt.Println(result.Path)
	// fmt.Println(result.Port())
	// fmt.Println(result.RawQuery)

	qparams := result.Query()

	fmt.Printf("the type of query params are: %T\n", qparams)

	fmt.Println(qparams["envId"])

	for _, val := range qparams {
		fmt.Println("param is : ", val)
	}

	// constructing url

	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=harshit",
	}

	anotherURL := partsOfUrl.String()

	fmt.Println(anotherURL)
}
