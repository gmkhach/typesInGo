package main

import "fmt"

// Here is a variable of type int
var x = 42

// Here is another variable of type string
var y = `"With the first link the chain is forged" - Aaron Satie`

// These variables are now locked to those types because GO is a static programming language
// You cannot assign anything but an int to x, and a string to y.

func main() {
	// Low format printing that shows the type of variable x
	// %T indicates that the type of the variable should be printed 
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
}