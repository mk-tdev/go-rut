package main

import (
	"fmt"
	"strings"
)

const defaultLength = 20

func printSeparator(length int) {
	if length <= 0 {
		length = defaultLength
	}
	fmt.Println(strings.Repeat("-", length))
}