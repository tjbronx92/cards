package main

func main() {
	cards := deck{"Ace of Diamonds", newCard()}
	cardx := append(cards, "Six of Spades")

	cardx.print()
}

func newCard() string {
	return "Five of Diamonds"
}
