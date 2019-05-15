package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {

	attacksAndCounters := map[rune]rune{
		'R': 'S',
		'B': 'K',
		'L': 'H',
	}

	reader := bufio.NewReader(os.Stdin)
	attacks, _ := reader.ReadString('\n')

	queue := []rune(attacks)
	queue = queue[:len(queue)-2] // remove the last two elements ('\r\n')
	var counters bytes.Buffer

	for len(queue) != 0 {
		if len(queue) >= 3 && queue[0] != queue[1] && queue[1] != queue[2] && queue[0] != queue[2] {
			queue = queue[3:]
			counters.WriteString("C")
		} else {
			counters.WriteString(string(attacksAndCounters[queue[0]]))
			queue = queue[1:]
		}
	}
	fmt.Print(counters.String())
}
