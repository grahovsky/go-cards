package main

import (
	"fmt"
)

func main() {

	card := newCard()
	fmt.Println(card)

	//printState()
	// go run main.go state.go

}

func newCard() string {
	return "Five of Diamonds"
}
