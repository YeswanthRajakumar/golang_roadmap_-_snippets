package main

import (
	"fmt"
	"time"
)

func printGreeting(b bot) {
	fmt.Println(b.greet())
}

type bot interface {
	greet() string
	getDate() time.Time
}
