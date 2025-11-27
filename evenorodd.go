package main

import "fmt"

type slice []int

func sliceOfNumbers() slice {
	return slice{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
}

func evenOrOdd(slice slice) {
	for index, num := range slice {
		if num%2 == 0 {
			fmt.Println(num, "is even in index", index)
		} else {
			fmt.Println(num, "is odd in index", index)
		}
	}
}

func (s slice) printEvenOrOdd() {
	for _, num := range s {
		if num%2 == 0 {
			fmt.Println(num, "is even")
		} else {
			fmt.Println(num, "is odd")
		}
	}
}