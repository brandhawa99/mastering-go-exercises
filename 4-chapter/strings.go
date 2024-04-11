package main

// PART 1:
import (
	"fmt"
)

/*
	Go strings are value types not pointers
	UTF-8 support by default
	Strings are read only byte slice
	- can hold any type of byte and can have an arbitrary length

*/

func main() {
	// PART 2:
	const sLiteral = "\x99\x42\x32\x55\x50\x35\x23\x50\x29\x9c" // each \xAB sequence represents a single char
	fmt.Println(sLiteral)
	fmt.Printf("x: %x\n", sLiteral)                    // %x returns the AB part of \xAB
	fmt.Printf("sLiteral length: %d\n", len(sLiteral)) // len(string) to find length

	// PART 3:
	for i := 0; i < len(sLiteral); i++ {
		fmt.Printf("%x ", sLiteral[i])
	}
	fmt.Println()

	fmt.Printf("q: %q\n", sLiteral)   // print a double-quoted string safely escaped
	fmt.Printf("+q: %+q\n", sLiteral) // guarantee ASCII-only output
	fmt.Printf("x: % x\n", sLiteral)  // % x*note the space*  will put spaces between the printed bytes

	fmt.Printf("s: As a string: %s\n", sLiteral) // prints a string literal

	// PART 4:
	s2 := "€£³"            // contains Unicode characters
	for x, y := range s2 { // using range allows you to process stings unicode characters one by one
		fmt.Printf("%#U starts at byte position %d\n", y, x) // %#U print the characters in U+0058 format
	}
	fmt.Printf("s2 length: %d\n", len(s2)) // len is 7 cuz they are unicode chars

	// Last Part
	const s3 = "ab12AB"
	fmt.Println("s3:", s3)
	fmt.Printf("x: % x\n", s3)

	fmt.Printf("s3 length: %d\n", len(s3))

	for i := 0; i < len(s3); i++ {
		fmt.Printf("%x ", s3[i])
	}
	fmt.Println()
}
