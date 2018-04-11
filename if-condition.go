package main

import "fmt"

func main() {

	incomingCall := "John"

	// assigning a value to the variable incomingCall

	if incomingCall == "John" {
		// conditioning
		fmt.Println("Attend the call")
	} else if incomingCall == "Arun" {
		// if else
		fmt.Println("Forward the call")
	} else {
		// last option else
		fmt.Println("Reject the call")
	}
}

// Note : you can use if, elseif, nested if with the same pattern
