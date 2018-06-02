package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func hasDuplicates(elements []string) bool {
	// Use map to record duplicates as we find them.
	encountered := map[string]bool{}

	for v := range elements {
		if encountered[elements[v]] == true {
			return true
		} else {
			// Record this element as an encountered element.
			encountered[elements[v]] = true
		}
	}
	// Found no duplicates in elements.
	return false
}

func countSuits(cards []string) map[string]int {
	// Suits order: P K H T
	suitCounter := map[string]int{
		"P": 0,
		"K": 0,
		"H": 0,
		"T": 0,
	}

	for i := range cards {
		switch string([]rune(cards[i])[0]) {
		case "P":
			suitCounter["P"]++
		case "K":
			suitCounter["K"]++
		case "H":
			suitCounter["H"]++
		case "T":
			suitCounter["T"]++
		default:
			fmt.Print(string([]rune(cards[i])[0]))
			fmt.Println(": Suit not found")
		}
	}

	return suitCounter
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	inputCards, _ := reader.ReadString('\n')

	cardFormat := "\\D\\d{2}"
	r, _ := regexp.Compile(cardFormat)
	cards := r.FindAllString(inputCards, -1)

	if hasDuplicates(cards) {
		fmt.Print("GRESKA")
	} else {
		maxSuits := 13
		suits := countSuits(cards)
		suitP := maxSuits - suits["P"]
		suitK := maxSuits - suits["K"]
		suitH := maxSuits - suits["H"]
		suitT := maxSuits - suits["T"]

		fmt.Printf("%v %v %v %v", suitP, suitK, suitH, suitT)
	}
}
