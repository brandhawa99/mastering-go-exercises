package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("this app keeps reading inputs until you type 'STOP'")

	for {
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		if text == "STOP" {
			os.Exit(1)
		}
		fmt.Println("you entered", text)

	}
}
