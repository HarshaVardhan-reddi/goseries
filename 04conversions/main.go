package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){
	fmt.Println("welcome to pizza app")
	fmt.Println("Please rate pizza btwn 1 to 5")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if(err != nil){
		fmt.Println("Somehitng is fisled")
	}

	fmt.Println("Thanks for rating",input)
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if(err != nil){
		fmt.Println("Failed to parse to float",err)
		panic(err) // gonna end programm here
	} else {
		fmt.Println("Added 1 to your rating: ", numRating+1)
	}
}