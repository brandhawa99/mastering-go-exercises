package main

import "fmt"

func main() {
	type XYZ struct {
		X int
		Y int
		Z int
	}
	p1 := XYZ{23, 12, -2}
	p2 := XYZ{X: 23, Z: 12}

	fmt.Println(p1)
	fmt.Println(p2)

}
