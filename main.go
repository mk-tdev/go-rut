package main

// func main() {
// 	var name string = "World"
// 	var age int = 25
// 	var isCool bool = true
// 	var size float64 = 25.5
// 	street := "Some street"
// 	name = "John"

// 	fmt.Println("Hello", name)
// 	fmt.Println("Age", age)
// 	fmt.Println("IsCool", isCool)
// 	fmt.Println("Size", size)
// 	fmt.Println("Street", street)
// }

// func main() {
// 	// card := newCard()
// 	// cards := []string{newCard(), "Two of Hearts"}
// 	// cards := deck{newCard(), "Two of Hearts"}

// 	cards := newDeck()
// 	// cards = append(cards, "Five of clubs")

// 	// fmt.Println(card)
// 	fmt.Println(cards)
// 	// fmt.Println(strings.Join(cards, ", "))

// 	cards.printDeck()
// }

// func main() {
// 	cards := newDeck()

// 	// fmt.Println(cards)
// 	cards.printDeck()

// 	printSeparator(50)

// 	dealed, _ := deal(cards, 5);
// 	dealed.printDeck()
// 	printSeparator(50)
// 	// remaining.printDeck();

// 	greeting := "Hello"
// 	fmt.Println("v", []byte(greeting))
// 	printSeparator(50)

// 	// fmt.Println("dealed", dealToString(dealed, 5))
// 	fmt.Println("dealed", dealed.dealToString())
// 	printSeparator(40)

// 	cards.saveToFile("defcards.txt")

// 	localCards := newDeckFromFile("cardss.txt", false)
// 	localCards.printDeck()
// }

func main() {
	cards := newDeck()
	cards.printDeck()
	printSeparator(50)

	cards.shuffleDeck()
	cards.printDeck()
	printSeparator(50)
}