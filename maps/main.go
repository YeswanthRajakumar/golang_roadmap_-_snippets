package main

import (
	"fmt"
)


func main() {
	// var colors map[string]string
	// colors := make(map[string]string)
	colors := map[string]string{
		"red":  "#0FFFDS",
		"blue": "#DDGFJ",
	}
	colors["green"] = "#FFFFF"

	delete(colors,"red")

	fmt.Println(colors)

	for k,v := range colors {
		fmt.Println(k," => ",v)
	}
}