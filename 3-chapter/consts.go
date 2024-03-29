package main

import "fmt"

type Power2 int

func main() {

	const (
		p2_0 Power2 = 1 << iota
		_
		p2_2
		_
		p2_4
		_
		p2_6
	)

	fmt.Println("p^0:", p2_0)
	fmt.Println("p^2:", p2_2)
	fmt.Println("p^4:", p2_4)
	fmt.Println("p^6:", p2_6)
}
