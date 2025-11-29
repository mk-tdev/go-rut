package main

import "fmt"


func main() {

	printMap(map[string]string{
		"red": "FF0000",
		"green": "00FF00",
		"blue": "0000FF",
	})
}

func printMap(c map[string]string) {

	fmt.Println(c)

	for color, hex := range c {
		fmt.Println(color, hex)
	}
}

// map vs struct
// map is a collection of key-value pairs
// struct is a collection of fields
// map is unordered
// struct is ordered
// map is mutable
// struct is immutable
// map is slower
// struct is faster
// map is easier to use
// struct is more powerful
// example:
// map[string]string
// struct {
// 	name string
// 	age int
// }

