package main

import (
	"fmt"
	"math/rand"
	// "time"
)

func main(){
	fmt.Println("")
	// rand.NewSource(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Printf("The diceNumber is %v\n",diceNumber)
	switch(diceNumber){
	case 1:
		fmt.Printf("Option selected is %v", diceNumber)
	case 2:
	fmt.Printf("Option selected is %v", diceNumber)
	case 3:
	fmt.Printf("Option selected is %v", diceNumber)
	case 4:
	fmt.Printf("Option selected is %v", diceNumber)
	fallthrough // Goes to the next block of statement. means executes case 5 too
	case 5:
	fmt.Printf("Option selected is %v", diceNumber)
	case 6:
	fmt.Printf("Hey you have one more chance %v", diceNumber)
	default:
		fmt.Println("Default: No option selected")
	}
}