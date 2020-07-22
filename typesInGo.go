package main

import "fmt"

// Here is a variable of type int
var x = 42

// Here is another variable of type string
var y = `"With the first link the chain is forged." - Aaron Satie`

// These variables are now locked to those types because GO is a static programming language
// You cannot assign anything but an int to x, and a string to y.

// One of the pillars of GO programming languages is the ability to declare your own types
// Here we are creating a type called awesome, with an underlyning type of int
type awesome int
var a awesome

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

	a = 7
	fmt.Printf("This is a: %v\tType: %T\n", a, a)

	// Although both the underlying type of awesome is int we cannot make assignments like x=a because x is of type int and a is of type awesome.
	// Instead we should convert a to type in first, and only then assign it to x. 
	// In GO programming language this is called conversion, not casting.
	x = int(a)
	fmt.Printf("This is x: %v\tType: %T\n", x, x)
}