package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGetting(eb)
	printGetting(sb)
}

func printGetting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	//VERY custom logic for generating an english greeting
	return "Hi There!"
}

func (spanishBot) getGreeting() string {
	//VERY custom logic for generating an english greeting
	return "Hola!"
}
