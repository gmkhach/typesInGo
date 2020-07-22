package main

import "fmt"

// Here is a variable of type int
var x = 42

// Here is another variable of type string
var y = `"With the first link the chain is forged" - Aaron Satie`

// These variables are now locked to those types because GO is a static programming language
// You cannot assign anything but an int to x, and a string to y.

func main() {
	// Using different verbs of format printing of fmt package
	fmt.Println(x)
	fmt.Printf("Type: %T\n", x)
	fmt.Printf("In decimal: %v\n", x)
	fmt.Printf("In binary: %#b\n", x)
	fmt.Printf("In hexadecimal: %X\n", x)
	fmt.Printf("In hexadecimal (go-syntax): %#X\n", x)

	// The same can be achieved in one Printf statement, but a value needs to be passed for each verb
	fmt.Println(y)
	fmt.Printf("Type: %T\nValue: %v\nQuoted: %q\nHexadecimal: %X\n", y, y, y, y )

	// You can also print into a string variable using Sprintf()
	s := fmt.Sprintf("Integer: %v\tstring: %v", x, y)
	fmt.Println(s)
}