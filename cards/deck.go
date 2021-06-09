package main

import "fmt"
// Something like class deck extends string slice

type deck []string

func (d deck) display()  {
	for _, card := range d {
		fmt.Println(card)
	}
}