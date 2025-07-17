package main

import (
	"fmt"
	"time"
)

func main(){
	fmt.Println("Welcome to time study of golang")
	presentTime := time.Now()
	fmt.Println(presentTime.Clock())
	fmt.Println(presentTime)
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	// creating date
	createdDate := time.Date(2020,time.June,12,12,10,12,12,time.Local)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01-02-2006 Monday"))

	// tickers

}

// Go build creates an executable file for the program
// go build 
// GOOS="windows" go build -> Creates executable for windows
// GOOS="linux" go build -> Creates executable for linux