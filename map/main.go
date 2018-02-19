package main

import (
	"fmt"
)

func main() {
	// var colors map[string]string]
	colors := make(map[string]string)
	colors["white"] = "#ffffff"
	colors["green"] = "#ffee22"
	colors["black"] = "#000000"
	// colors := map[string]string{
	// 	"red":   "#ff0011",
	// 	"green": "#ff0111",
	// }
	// delete(colors, "black")
	printMap(colors)
}

func printMap(m map[string]string) {
	for color, hex := range m {
		fmt.Println("Color is: ", color, " Hex Code is: ", hex)
	}
}
