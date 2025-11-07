package locksingo

import (
	"fmt"
	"sync"
)

var mutexLock sync.Mutex
var wg sync.WaitGroup

var GlobalValue = 5
func ExecutionOfThreadsWithLock(){
	tmp := GlobalValue
	defer wg.Wait()
	wg.Add(2) // we have two go routines with race conditon handled
	go addOneToGV()
	go subsOneToGV()

	wg.Wait()

	fmt.Printf("\nGlobal value before modification %d and Global value after modification %d", tmp, GlobalValue)
}

func addOneToGV(){
	defer wg.Done()
	fmt.Println("Adding 1 to global value")
	mutexLock.Lock()
	GlobalValue++
	mutexLock.Unlock()
}


func subsOneToGV(){
	defer wg.Done()
	fmt.Println("subsctrating 1 from global value")
	mutexLock.Lock()
	GlobalValue--
	mutexLock.Unlock()
}