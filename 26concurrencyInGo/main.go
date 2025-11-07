package main

import (
	"concurrencyInGo/locksingo"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
const NoOfGoRoutines = 5

func main(){
	fmt.Println("Welcome to concurrency in go")
	simpleConcurrentOperation()

	invokingAPoolOfThreads(NoOfGoRoutines)

	fmt.Println("🔒 Lock mechanisms in golang using mutex")
	wg.Wait() // waits until all are done

	locksingo.ExecutionOfThreadsWithLock()
}

func simpleConcurrentOperation(){
	go fmt.Println("Printing the background thread") // Go keywords performs the block of code in a different thread
	fmt.Println("printing the foreground thread")
	time.Sleep(2*time.Second)
}

func invokingAPoolOfThreads(lenofloop int){

	fmt.Println("\nExecuting threads Parallely 🚀")

	wg.Add(lenofloop) // how many go routines will be runned, that many
	for i := 0; i < lenofloop; i++{
		go func ()  {
			fmt.Println("Printing the current index", i)
			wg.Done() // once run is finished threads says i am done using the weight group function
		}()
	}
}