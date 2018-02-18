package main

import "fmt"

func main() {
	fmt.Println("Welcome to Go")
	cards := newDeck()
	hand, remaining := deal(cards, 5)
	hand.print()
	fmt.Println("============================")
	remaining.print()

}
