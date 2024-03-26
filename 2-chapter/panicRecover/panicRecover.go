package main

import "fmt"

func a() {
	fmt.Println("Inside a()")
	defer func() {
		if c := recover(); c != nil {
			fmt.Println("Recover inside a()!")
		}
	}()
	fmt.Println("About to call b()")
	b()
	fmt.Println("b() exited")
	fmt.Println("Exiting a()")
}

func b() {
	fmt.Println("Inside b()")
	panic("panic in b()!")
	fmt.Println("Exiting b()") // show that we don't get here
}

func main() {
	a()
	fmt.Println("main() ended")
}
