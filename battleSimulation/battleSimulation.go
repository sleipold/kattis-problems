package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func defend(offSeq string, permutations []string) string {
	defSeq := ""
	for _, char := range offSeq {
		switch attack := char; attack {
		case 'R':
			defSeq += "S"
		case 'B':
			defSeq += "K"
		case 'L':
			defSeq += "H"
		}
	}

	for _, currPerm := range permutations {
		if idx := strings.Index(defSeq, currPerm); idx != -1 {
			defSeq = strings.Replace(defSeq, currPerm, "C", 1)
		}
	}

	return defSeq
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	offSeq, _ := reader.ReadString('\n')
	permutations := make([]string, 0)
	permutations = append(permutations, "SKH", "KSH", "KHS", "SHK", "HSK", "HKS")
	defSeq := defend(offSeq, permutations)
	fmt.Println(defSeq)
}
