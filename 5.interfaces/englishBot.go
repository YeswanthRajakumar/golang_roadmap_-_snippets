package main

import "time"

type englishBot struct{}

func (en englishBot) getDate() time.Time {
	return time.Now()
}

func (en englishBot) greet() string {
	return "Hello World"
}
