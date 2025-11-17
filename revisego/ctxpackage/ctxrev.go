package ctxpackage

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var ctx context.Context

func BasicChannel(){
	fmt.Println("Basic channel printing")
	chn := make(chan string, 1)
	chn <- "Testing somehting"
	// time.Sleep(2 * time.Millisecond)
	println(<-chn)
}


func TestingGoRoutines(){
	ctx = context.Background()
	ctx, cancel := context.WithTimeout(ctx, 1000 * time.Millisecond)

	defer cancel()
	var wg sync.WaitGroup
	wg.Add(2)
	chn := make(chan string, 1)
	go func(){
		fmt.Println("Sending message into the channels:")
		chn<- "This is the message from channels.."
		wg.Done()
	}()
	go func(){
		fmt.Println("Reading channel: ", <-chn)
		wg.Done()
	}()
	time.Sleep(1200 * time.Millisecond)
	select {
	case <- ctx.Done():
		fmt.Println(ctx.Err().Error())
	case res := <-chn:
		fmt.Println("This is the resl go from the other channesl", res)
	// default:
	// 	fmt.Println("Done with out timeout")
	}
	wg.Wait()
}