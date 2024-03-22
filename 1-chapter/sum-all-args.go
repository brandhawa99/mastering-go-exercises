package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	var sum int = 0

	args := os.Args

	for i := 1; i <= len(args)-1; i++ {
		num, _ := strconv.Atoi(args[i])
		sum += num

	}

	fmt.Println("sum of all the values is", sum)
}
