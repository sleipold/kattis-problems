package main

import (
	"bufio"
	"fmt"
	"os"
)

func mapCards(inputCards string) map[string]int {
	cards := map[string]int{
		"T": 0,
		"C": 0,
		"G": 0,
	}

	for _, card := range inputCards {
		switch card {
		case 'T':
			cards["T"]++
		case 'C':
			cards["C"]++
		case 'G':
			cards["G"]++
		}
	}

	return cards
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func calcScore(cards map[string]int) int {
	score := 0
	for i := range cards {
		score += cards[i] * cards[i]
	}
	score += min(min(cards["T"], cards["C"]), cards["G"]) * 7
	return score
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	inputCards, _ := reader.ReadString('\n')

	cards := mapCards(inputCards)
	score := calcScore(cards)
	fmt.Println(score)
}
