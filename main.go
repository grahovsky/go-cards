package main

import "log"

func main() {

	cards := newDeck()

	hand, remainingDeck := deal(cards, 5)

	hand.print()
	remainingDeck.print()

	err := hand.saveToFile("saveData")
	if err != nil {
		log.Fatal(err)
	}

}
