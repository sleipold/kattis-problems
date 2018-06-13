package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func rotateCW(mat [4][4]int) [4][4]int {
	rotMat := [4][4]int{}

	for r := 0; r < len(mat); r++ {
		for c := 0; c < len(mat[0]); c++ {
			rotMat[c][len(mat)-1-r] = mat[r][c]
		}
	}

	return rotMat
}

func execMove(board [4][4]int, direction int) [4][4]int {
	movedBoard := board

	for i := 0; i < 4-direction; i++ {
		movedBoard = rotateCW(movedBoard)
	}

	for row := 0; row < 4; row++ {
		for curr := 0; curr < 4; curr++ {
			for next := curr + 1; next < 4; next++ {
				if movedBoard[row][next] == 0 {
					// do nothing, check next -> next++
				} else if movedBoard[row][curr] == movedBoard[row][next] {
					// same value -> merge
					movedBoard[row][curr] *= 2
					movedBoard[row][next] = 0
					break
				} else if movedBoard[row][curr] == 0 {
					// current is zero -> move next left
					movedBoard[row][curr] = movedBoard[row][next]
					movedBoard[row][next] = 0
				} else if next != curr+1 {
					// next not directly next to curr -> zero between curr and next -> move next left
					movedBoard[row][curr+1] = movedBoard[row][next]
					movedBoard[row][next] = 0
					break
				} else {
					break
				}
			}
		}
	}

	for i := 0; i < direction; i++ {
		movedBoard = rotateCW(movedBoard)
	}

	return movedBoard
}

func numbers(s string) []int {
	var n []int
	for _, f := range strings.Fields(s) {
		i, err := strconv.Atoi(f)
		if err == nil {
			n = append(n, i)
		}
	}
	return n
}

func readInput() ([4][4]int, int) {
	board := [4][4]int{}
	direction := 0

	var firstLine, secondLine, thirdLine, fourthLine []int
	scanner := bufio.NewScanner(os.Stdin)
	for i := 1; i <= 4 && scanner.Scan(); i++ {
		switch i {
		case 1:
			firstLine = numbers(scanner.Text())
		case 2:
			secondLine = numbers(scanner.Text())
		case 3:
			thirdLine = numbers(scanner.Text())
		case 4:
			fourthLine = numbers(scanner.Text())
		}
	}

	fmt.Scanf("%d", &direction)

	for i := 0; i < 4; i++ {
		board[0][i] = firstLine[i]
	}

	for i := 0; i < 4; i++ {
		board[1][i] = secondLine[i]
	}

	for i := 0; i < 4; i++ {
		board[2][i] = thirdLine[i]
	}

	for i := 0; i < 4; i++ {
		board[3][i] = fourthLine[i]
	}

	return board, direction
}

func printBoard(board [4][4]int) {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			fmt.Print(board[i][j])
			fmt.Print(" ")
		}
		fmt.Print("\n")
	}
}

func main() {
	board, direction := readInput()
	board = execMove(board, direction)
	printBoard(board)
}
