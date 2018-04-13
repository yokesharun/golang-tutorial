package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName  string
}

// delcare the structs with no comma's

func assigning() {

	arun := person{"Arun", "Yokesh"}

	arun.PrintReceiver()
	// to call the receiver function
}

// receiver function are like a protected function it will receive only the declared type
// here we are restricted to receive only the type of person and print them

func (r person) PrintReceiver() {
	fmt.Println(r)
}
