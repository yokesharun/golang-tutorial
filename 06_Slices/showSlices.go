package main

import "fmt"

func showslices() {

	// Slices are a key data type in Go, giving a more powerful interface to sequences than arrays.

	// An array thet can grow or shrink

	cards := []string{"a", "b"}

	cards = append(cards, "c")
	cards = append(cards, "d")
	// appending new elemet to the slice

	// printing all slice elements

	for _, card := range cards {
		fmt.Println(card)
		// printing each card value
	}

	// slice[startIndesIncluding: uptoNotIncluding]

	fmt.Println(cards[2])
	// OUTPUT : "c"

	fmt.Println(cards[0:2])
	// OUTPUT : "a", "b", "c"

	fmt.Println(cards[:2])
	// OUTPUT : "a", "b"

	fmt.Println(cards[2:])
	// OUTPUT : "c", "d"
}
