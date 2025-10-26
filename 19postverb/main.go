package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main(){
	fmt.Println("sending a post request in golang")
	PerformPostRequest()
}

func PerformPostRequest(){
	const myurl = "https://api.restful-api.dev/objects"
	reuqestBody := strings.NewReader(`
	{
   "name": "Apple MacBook Pro 16",
   "data": {
      "year": 2019,
      "price": 1849.99,
      "CPU model": "Intel Core i9",
      "Hard disk size": "1 TB"
   }
}
	`)

	response, err := http.Post(myurl, "application/json", reuqestBody)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)
	fmt.Println(string(content))
	
}