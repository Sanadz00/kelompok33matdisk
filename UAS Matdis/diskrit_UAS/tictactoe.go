package main

import (
	"fmt"
	"math"
)

var board = [9]string{" ", " ", " ", " ", " ", " ", " ", " ", " "}

func printBoard() {
	fmt.Println()
	fmt.Println(board[0], "|", board[1], "|", board[2])
	fmt.Println("--+---+--")
	fmt.Println(board[3], "|", board[4], "|", board[5])
	fmt.Println("--+---+--")
	fmt.Println(board[6], "|", board[7], "|", board[8])
	fmt.Println()
}

func checkWinner() string {
	winPatterns := [][]int{
		{0, 1, 2}, {3, 4, 5}, {6, 7, 8},
		{0, 3, 6}, {1, 4, 7}, {2, 5, 8},
		{0, 4, 8}, {2, 4, 6},
	}

	for _, p := range winPatterns {
		if board[p[0]] != " " &&
			board[p[0]] == board[p[1]] &&
			board[p[1]] == board[p[2]] {
			return board[p[0]]
		}
	}

	for _, v := range board {
		if v == " " {
			return ""
		}
	}

	return "draw"
}

func minimax(isMax bool) int {
	result := checkWinner()
	if result == "O" {
		return 1
	} else if result == "X" {
		return -1
	} else if result == "draw" {
		return 0
	}

	if isMax {
		best := math.MinInt32
		for i := 0; i < 9; i++ {
			if board[i] == " " {
				board[i] = "O"
				score := minimax(false)
				board[i] = " "
				if score > best {
					best = score
				}
			}
		}
		return best
	} else {
		best := math.MaxInt32
		for i := 0; i < 9; i++ {
			if board[i] == " " {
				board[i] = "X"
				score := minimax(true)
				board[i] = " "
				if score < best {
					best = score
				}
			}
		}
		return best
	}
}

func bestMove() {
	bestScore := math.MinInt32
	move := -1

	for i := 0; i < 9; i++ {
		if board[i] == " " {
			board[i] = "O"
			score := minimax(false)
			board[i] = " "
			if score > bestScore {
				bestScore = score
				move = i
			}
		}
	}
	if move != -1 {
		board[move] = "O"
	}
}

func main() {
	for {
		printBoard()

		var pos int
		fmt.Print("Masukkan posisi (1-9): ")
		_, err := fmt.Scan(&pos)

		if err != nil || pos < 1 || pos > 9 {
			fmt.Println("Input tidak valid! Masukkan angka 1-9")
			continue
		}

		pos--

		if board[pos] != " " {
			fmt.Println("Posisi sudah terisi!")
			continue
		}

		board[pos] = "X"

		if checkWinner() != "" {
			printBoard()
			break
		}

		// Periksa apakah masih ada posisi kosong
		hasMove := false
		for i := 0; i < 9; i++ {
			if board[i] == " " {
				hasMove = true
				break
			}
		}

		if hasMove {
			bestMove()
		}

		if checkWinner() != "" {
			printBoard()
			break
		}
	}

	winner := checkWinner()
	if winner == "draw" {
		fmt.Println("Hasil: Seri")
	} else {
		fmt.Println("Pemenang:", winner)
	}
}
