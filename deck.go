package main

import "fmt"

type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println("Index: ", i, " Card: ", card)
	}
}

func newDeck() deck {
	cards := deck{}
	cardSuits := [4]string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := [4]string{"Ace", "Two", "Three", "Four"}
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}
