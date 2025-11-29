package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
}

func main() {
	p := Person{
		firstName: "Jane",
		lastName:  "Doe",
	}
	p.firstName = "John"
	p.lastName = "Doe"

	var p2 Person

	fmt.Println(p)
	fmt.Printf("%+v\n", p)
	fmt.Println(p2)

	fmt.Println("test")

}
