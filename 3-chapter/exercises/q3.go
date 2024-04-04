package main

import "fmt"

func main() {

	var arr = [3]string{"1", "2", "3"}

	m := make(map[int]string)
	for idx, value := range arr {
		m[idx] = value
	}

	for key, value := range m {
		fmt.Printf("Map Key: %d value %s \n", key, value)
	}
}
