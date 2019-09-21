package main

import "fmt"

func main() {
	var mb, months, mbUsedInput int
	fmt.Scan(&mb, &months)

	mbUsedSum := 0
	for i := 0; i < months; i++ {
		fmt.Scan(&mbUsedInput)
		mbUsedSum += mbUsedInput
	}

	fmt.Print((mb * months) - mbUsedSum + mb)
}
