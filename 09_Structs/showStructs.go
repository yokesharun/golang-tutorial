package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName  string
}

// Go’s structs are typed collections of fields. They’re useful for grouping data together to form records.

// delcare the structs with no comma's

func showStruct() {

	arun := person{"Arun", "Yokesh"}

	// or

	Language := person{firstName: "Go", lastName: "Lang"}

	fmt.Println(arun)
	fmt.Println(Language)

	// structs can be updated only with pointers
}
