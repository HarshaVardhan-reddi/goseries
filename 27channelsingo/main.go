package main

import (
	"fmt"
	"sync"
)

func main(){
	fmt.Println("Welcome to channels & deadlocks in golang")
	// accessing the common memory, SHARED MEMORY, SIGNALS e.t.c
	// channels will be used for inter process communication in golang
	// multiple go routine can talk to each other
	// chan is the keyword | More of like a pipe btwn two go routines
	// myChannel := make(chan int)
	// myChannel <- 1 // this is how values are pushed into channels | Always arrow points towards the channel
	// fmt.Println(<-myChannel) // this is how we access values in channel

	// we can push values to channel if somebody is listining or else we not pass | the above will run into a deadlock util line 15
	myChannel := make(chan int, 2) // buffer channel
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func (ch <-chan int, wg *sync.WaitGroup){ // recieve only | if recive only channel we cannot even close the channel
		val, isChannelOpen := <-ch
		if isChannelOpen{
			fmt.Println("Printing the value from channel", val) // its starts listening here from the channel, this prints 5
			fmt.Println("Printing the value from channel", <-ch)// if there are two publishers then there should be two listerners, as passing 5&6 so processing again. this will print 6
			// close(ch)
		}
		defer wg.Done()
	}(myChannel,wg)

	go func (ch chan<- int, wg *sync.WaitGroup){ // send only 
		fmt.Println("Pushing value 5 in to the channel")
		ch <- 5
		ch <- 6
		defer wg.Done()
	}(myChannel,wg)

	wg.Wait()
}