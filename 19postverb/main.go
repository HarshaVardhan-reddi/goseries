package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type DataStrct struct {
	Year int
	Price float32
	CPUModel string `json:"CPU model"`
	HardDiskSize string `json:"Hard disk size"`
}

type Info struct{
	Id string
	Name string
	Data DataStrct
}

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

	info := Info{}

	errParse := json.Unmarshal(content, &info)
	if errParse != nil {
		panic(errParse)
	}

	fmt.Println(info)
	
}