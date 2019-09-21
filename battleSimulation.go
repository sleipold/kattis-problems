package main

import (
	"bufio"
	"fmt"
	"os"
)

func defend(offSeq string) string {
	defSeq := ""

	for i := 0; i < len(offSeq); i++ {
		if len(offSeq)-i > 3 && offSeq[i]+offSeq[i+1]+offSeq[i+2] == 224 {
			defSeq += "C"
			i += 2
		} else {
			switch string(offSeq[i]) {
			case "R":
				defSeq += "S"
			case "B":
				defSeq += "K"
			case "L":
				defSeq += "H"
			}
		}
	}

	return defSeq
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	offSeq, _ := reader.ReadString('\n')

	fmt.Println(defend(offSeq))
}
