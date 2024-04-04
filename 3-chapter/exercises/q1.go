package main

import "fmt"

func main() {

	const (
		p4_0 int = 1 << (iota * 2)
		p4_2
		p4_4
		p4_5
	)
	fmt.Println(p4_0)
	fmt.Println(p4_2)
	fmt.Println(p4_4)
	fmt.Println(p4_5)
}
