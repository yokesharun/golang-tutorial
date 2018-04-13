package main

import "fmt"

func main() {
	colors := make(map[string]string)
	// [string] -> this is the key data type
	// string -> value datatype

	// adding a value to map

	colors["red"] = "#ff0000"
	colors["black"] = "#000000"

	// deleting a map value

	delete(colors, "red")

	// printing a map values

	for color, hexa := range colors {
		fmt.Println(hexa)
	}
}
