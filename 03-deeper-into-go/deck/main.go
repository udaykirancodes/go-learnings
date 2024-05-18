package main

import (
	"fmt"
)

func main(){

	cards := newDeck()

	// cards := newDeckFromFile("my_cards.txt");

	cards.shuffle();
	// cards.saveToFile("my_cards.txt");

	fmt.Println(cards.toString());


	// first , second := deal(cards, 20)

	// first.print();
	// fmt.Println()
	// fmt.Println()
	// second.print();

}

