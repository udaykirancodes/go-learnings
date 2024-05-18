package main

import "fmt"

func main(){
	// Slice of Cards 
	cards := []string{
		"Ace of Diamonds",newCard(),
	}

	// append 
	cards = append(cards, "Six of Spades")
	fmt.Println(cards);

	// Printing cards
	for i,card := range cards{
		fmt.Println(i,card);
	}
}
func newCard() string{
	return "Five of Diamonds"
}