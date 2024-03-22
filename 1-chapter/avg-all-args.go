package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	//
	args := os.Args
	var total = 0
	if len(args) <= 1 {
		fmt.Println("No arguments passed")
		os.Exit(1)

	} else {
		for i := 1; i < len(args); i++ {
			num, _ := strconv.Atoi(args[i])
			total += num
		}
	}
	avg := float32(total) / float32(len(args)-1)
	fmt.Printf("the average of all the values was %.2f, the total was %d \n", avg, total)
}
