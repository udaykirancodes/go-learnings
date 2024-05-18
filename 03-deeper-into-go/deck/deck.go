package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of 'deck'
// which is  a slice of strings

type deck []string 

func newDeck() deck{
	cards := deck{}

	cardSuits := []string{"Spades","Diamonds","Hearts","Clubs"}

	cardValues := []string{"Ace","Two","Three","Four"}

	for _, suit := range cardSuits{
		for _,value := range cardValues{
			cards = append(cards, value + " of " + suit);
		}
	}

	return cards
}

func deal (d deck , handSize int) (deck,deck) {
	return d[:handSize] , d[handSize :]
}

func  (d deck) toString() string {
 	return strings.Join([]string(d),",")
}

func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename , []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {

	bs , err := os.ReadFile(filename);

	if err != nil {
		fmt.Println("Error : ",err);	
		os.Exit(1); 
	}

	bsStr := string(bs);

	val := strings.Split( bsStr, ",");

	return deck(val)
}

func (d deck) shuffle() deck {
	
	for i := range d {
		rand.Seed(time.Now().UnixNano())
		newPosition := (rand.Intn(len(d) - 1))
		d[i] , d[newPosition] = d[newPosition] , d[i]
	}

	return deck{}
}

// this will add a function to the variables which are of type `deck`
func (d deck) print(){
	for i, card := range  d {
		fmt.Println(i,card);
	}
}

func (d deck) printF(){
	for i , card := range d {
		fmt.Println(i , " : ",card)
	}
}

