package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type ResponseStrcuture struct {
	Message map[string][]string
	Status string
}

func main(){
	fmt.Println("Welcome to web verb")
	PerformGetRequest()
}


func PerformGetRequest(){
	const myurl = "https://dog.ceo/api/breeds/list/all"
	response, err := http.Get(myurl)
	if err != nil{
		panic(err)
	}
	defer response.Body.Close()
	fmt.Println("Status code:", response.StatusCode)
	fmt.Println("Content length :", response.ContentLength)

	var responseString strings.Builder

	content, _ := io.ReadAll(response.Body)


	// byteCount, _ := responseString.Write(content)

	// fmt.Println(byteCount)
	// fmt.Println(responseString.String())
	fmt.Println(responseString.Len())
	responseStrct := ResponseStrcuture{}

	errl := json.Unmarshal(content, &responseStrct)

	if errl != nil {
		panic(errl)
	}

	// fmt.Println(responseStrct)
	// fmt.Println("Message: ",responseStrct.Message)
	fmt.Println("Status: ",responseStrct.Status)

	dogs := responseStrct.Message
	for name, behavoiur := range(dogs) {
		fmt.Println("Type of the dog is: ",name)
		fmt.Println("behavious of the dog: ", behavoiur)
	}

	// json.Decoder()

	// fmt.Println("printing the content: ",string(content))
}