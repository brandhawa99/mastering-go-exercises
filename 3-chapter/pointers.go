package main

import "fmt"

func main() {
	i := 2
	fmt.Println(&i)
}

func getPointer(n *int) {
	*n = *n * *n
}

func returnPointer(n int) *int {
	v := n * n
	return &v
}
