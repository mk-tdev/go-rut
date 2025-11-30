package main

import "fmt"

// A struct is a collection of fields.
// It is a way to define a data structure with multiple fields.
// It is a way to define a type that can hold multiple values.
// It is a way to define a type that can hold named values.
// It is a way to define a type that can hold multiple named values.
// It is a way to define a type that can hold multiple values of different types.

type englishBot struct {}
type spanishBot struct {}


// An interface is a set of methods that a type must implement in order to be considered to be that interface type.
// Interface is a contract that a type must follow.
// It is a way to define a set of methods that a type should have.
// It is a way to specify what a type should look like.

type bot interface {
	getGreeting() string
	getBotType() string
	getBotVersion() float64
}

func main() {
	fmt.Println("Interfaces Rut")

	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getBotType(), " says: ", b.getGreeting(), " Version: ", b.getBotVersion())
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

func (eb englishBot) getBotVersion() float64 {
	return 1.0
}

func (sb spanishBot) getBotVersion() float64 {
	return 2.0
}