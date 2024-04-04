package main

import "fmt"

func main() {

	const (
		MONDAY = iota
		TUESDAY
		WEDNESDAY
		THURSDAY
		FRIDAY
		SATURDAY
		SUNDAY
	)

	fmt.Println(MONDAY)
	fmt.Println(TUESDAY)
	fmt.Println(WEDNESDAY)

}
