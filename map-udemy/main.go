package main

import "fmt"

func main(){
	colors := map[string]string{
		"red": "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	// var colors map[string]string

	// colors := make(map[string]string)

	// colors["white"] = "#ffffff"

	colors["yellow"] = "#ywywyw"


	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("hex code for", color, "is", hex)
	}

}