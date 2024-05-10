package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

var passOptions = []string{"pass", "go", "stay", "password", "lolol", "nvm"}

func main() {
	MIN := 0
	MAX := 94
	var LENGTH int64 = 8
	arguments := os.Args
	flag := false
	switch len(arguments) {
	case 2:
		LENGTH, _ = strconv.ParseInt(os.Args[1], 10, 64)
	case 3:
		flag = true
	default:
		fmt.Println("Using default values!")
	}

	startChar := "!"
	var i int64 = 1

	if !flag {

		for {
			myRand := random(MIN, MAX)
			newChar := string(startChar[0] + byte(myRand))
			fmt.Print(newChar)
			if i == LENGTH {
				break
			}
			i++
		}
	} else {
		MAX = len(passOptions)
		MIN = 0
		item := passOptions[random(MIN, MAX)]
		fmt.Print(item + "" + strconv.FormatInt(time.Now().Unix(), 10))

	}
	fmt.Println()
}
