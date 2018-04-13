package main

import "fmt"

func main() {
	// creating a function
	fmt.Println("I'm Main function")

	// calling another function
	child()

	// getting a return type of PrintInteger and assigning to ImInteger variable
	ImInteger := PrintInteger(1)

	fmt.Println(ImInteger)

}

func child() {
	fmt.Println("I'm Child function")
}

// function with parameter passing and with return type

func PrintInteger(num int) int {
	return num
}
