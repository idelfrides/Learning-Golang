package main

func main() {
	cards := newDeck()
	cards.shuffle()
	cards.ijprint()
}

/*
filename := "my_cards_v2"
	// filenameTest := "my_cards_not"

	cards := newDeck()

	cards.ijprint()
	hand, remainingCards := deal(cards, 5)
	hand.ijprint()
	remainingCards.ijprint()

	// fmt.Println(cards.toString())
	cards.saveToFile(filename)

	// cardsFromdisk := newDeckFromFile(filenameTest)
	cardsFromdisk := newDeckFromFile(filename)
	cardsFromdisk.ijprint()
*/
