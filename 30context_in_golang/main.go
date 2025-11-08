package main

import (
	"context"
	"fmt"
	"strconv"
	"sync"
	"time"
)

var ctx = context.Background();  // this is a non nil emprty context, this basically a empty struct

func main(){
	fmt.Println("Welcome to context in golang")
	simpleTimeOutUseCase()
	storingInfoInCtxUseCase()
	// killGoRoutinesWhichAreRunning()
}


func simpleTimeOutUseCase(){
	ctx, cancel := context.WithTimeout(ctx,5*time.Millisecond)
	//channel
	stream := make(chan string, 5)

	// done channel
	done := make(chan bool)
	defer cancel()


	go func(){
		fmt.Println("Inside anonymous function")
		for i:=0; i < 5; i++{
			stream<-strconv.Itoa(i)
		}
		time.Sleep(5 * time.Millisecond)
		close(done)
	}()

	for{
		select {
		case <-done:
			return
		case streamval := <-stream:
			fmt.Println("Here is where i am printing the val", streamval)
		case <-ctx.Done(): // if cxt time out happens then close(ctx.Done()) will be triggered basically. 
			fmt.Println("Sometgin is wrong took more time to process",ctx.Err().Error())
		}
	}
}


func storingInfoInCtxUseCase(){
	wg := sync.WaitGroup{}
	type number int
	num := number(1)
	wg.Add(1)

	type str string
	strval := str("UserId")

	ctx := context.WithValue(ctx,strval,num)
	// go routine
	go func(){
		fmt.Println("Accessing ctx inside the go routine")
		userId := ctx.Value(str("UserId")) // accessing the context key inside go routine
		if userId != nil{
			fmt.Println("user id is", userId)
		}
		wg.Done()
	}()

	wg.Wait()
}


// func killGoRoutinesWhichAreRunning(){
// 	ctx, cancel := context.WithCancel(ctx)
// 	go func(){
		
// 	}
// 	defer cancel()
// }