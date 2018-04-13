package main

// main package is an executable one, the name except the word main are the custom packages which is not executable

import "fmt"

// fmt is a standard Golang package which is commonly used.
// for more info ref : https://golang.org/pkg/fmt/

// every main file package must have the function called main, which executes when the package main calls

func main() {
	// creating a new function called main, to create a new function the key word is "func"

	fmt.Println("Hi There!")

	// for printing the text we are calling Println function from the package "fmt"
}

// Tips : Go is one of the Static Types Programming Language like C++ and Java
