package main

import "fmt"
type englishBot struct {}
type spanishBot struct {}
type bot interface {
	getGreeting() string
	getBotType() string
}

func main() {
	fmt.Println("Interfaces Rut")

	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getBotType(), " says: ", b.getGreeting())
}

func (eb englishBot) getGreeting() string {
	return "Hello from English"
}

func (sb spanishBot) getGreeting() string {
	return "Hola from Spanish"
}

func (eb englishBot) getBotType() string {
	return "English"
}

func (sb spanishBot) getBotType() string {
	return "Spanish"
}