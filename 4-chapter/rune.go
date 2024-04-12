package main

import "fmt"

/*
	rune is a int32 value therefore it's a Go type that is used for representing a Unicode code point
	- a numerical value that is used for representing a single Unicode character
	*rune literal* char in single quotes or rune constant
*/

func main() {

	const r1 = '€'
	fmt.Println("(int32) r1:", r1)              // gets printed as int32 value
	fmt.Printf("(HEX) r1: %x\n", r1)            // prints HEX value
	fmt.Printf("(as a string) r1: %s\n", r1)    // prints as string
	fmt.Printf("(as a character) r1: %c\n", r1) //prints as character

	fmt.Println("A string is a collection of runes:", []byte("Bawa"))
	aString := []byte("Bawa")
	for x, y := range aString {
		fmt.Println(x, y)
		fmt.Printf("Char: %c\n", aString[x])
	}
	fmt.Printf("%s\n", aString)

	/*
		get an illegal run literal error msg by using single quotes instead of double quotes when importing a package
	*/

}
