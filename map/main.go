package main

import (
	"fmt"
)

func main() {
	// map declaration ways
	// var mymap map[string]string
	// mymap2 := make(map[string]string)

	colors := map[string]string{
		"red":    "#ff0000",
		"green":  "#00ff00",
		"yellow": "#ffff66",
	}

	ages := map[string]int{
		"BRASIL": 201,
		"USA":    247,
		"RGB":    50,
	}

	colors["white"] = "#ffffff"
	colors["black"] = "#000000"

	printMap(colors)
	// printMap(ages)

	// fmt.Println(colors)
	fmt.Println(ages)
	// fmt.Printf("MY map --> %v", mymap)
	// fmt.Printf("\n MY map 2 --> %v", mymap2)

	delete(colors, "red")
	fmt.Println("\n\n AFTER DELETION")
	printMap(colors)
}

func printMap(c map[string]string) {

	for k, v := range c {
		fmt.Printf("\n %v = %v", k, v)
	}
}
