package tic_tac_toe

import (
	"fmt"
	"math/rand"
)

func IsGameOver(board [][]string) (bool, string) {
	for i := 0; i < 3; i++ {
		if board[i][0] == board[i][1] && board[i][0] == board[i][2] && board[i][0] != "" {
			return true, board[i][0]
		}
	}
	for i := 0; i < 3; i++ {
		if board[0][i] == board[1][i] && board[0][i] == board[2][i] && board[0][i] != "" {
			return true, board[0][i]
		}
	}
	if board[0][0] == board[1][1] && board[0][0] == board[2][2] && board[0][0] != "" {
		return true, board[0][0]
	}

	if board[0][2] == board[1][1] && board[0][2] == board[2][0] && board[1][1] != "" {
		return true, board[0][2]
	}

	return false, ""
}

func MakeMove(i, j int8, board [][]string, player string) ([][]string, bool) {
	moved := false
	if i < 0 || j < 0 || i > 2 || j > 2 {
		fmt.Println("Please enter a valid move")
	} else if board[i][j] != "" {
		fmt.Println("Please enter a valid move")
	} else {
		board[i][j] = player
		moved = true
		fmt.Println("Player ", player, " move", i, j)
		PrintBoard(board)
	}

	return board, moved
}

func PrintBoard(board [][]string) {
	for i := 0; i < 3; i++ {
		fmt.Print("| ")
		for j := 0; j < 3; j++ {
			if board[i][j] == "" {
				fmt.Print(" ")
			}
			fmt.Print(board[i][j], " | ")
		}
		fmt.Println()
		fmt.Print("-------------")
		fmt.Println()
	}
}

func InitializeGame() {
	board := make([][]string, 3)
	for i := range board {
		board[i] = []string{"", "", ""}
	}
	// board := [][]int8{{-1, -1, -1}, {-1, -1, -1}, {-1, -1, -1}}
	fmt.Print("Enter text: ")
	var i, j int8
	//fmt.Scanf("%d %d", &i, &j)
	var human, computer string = "X", "O"
	var moved bool
	prevPlayer := computer
	for a := 0; a < 9; {
		gameOver, player := IsGameOver(board)
		if gameOver {
			fmt.Println("Game over, Player", player, "wins")
			break
		}
		if prevPlayer == human {
			fmt.Println("Please enter next move in format x y. E.g. 1 2:")
			fmt.Scanf("%d %d", &i, &j)
			board, moved = MakeMove(i, j, board, human)
		} else {
			i, j = RandomCoordinates(board)
			board, moved = MakeMove(i, j, board, computer)
		}
		if moved {
			prevPlayer = ChangePlayer(prevPlayer)
			a++
		}
	}
}

func ChangePlayer(prevPlayer string) string {
	if prevPlayer == "O" {
		return "X"
	} else if prevPlayer == "X" {
		return "O"
	} else {
		panic("Not a valid player")
	}
}

func RandomCoordinates(board [][]string) (int8, int8) {
	a, b := GetRemainingSlots(board)
	randVal := rand.Intn(int(b))
	return a[randVal][0], a[randVal][1]
}

func GetRemainingSlots(board [][]string) ([][]int8, int8) {
	remaining := make([][]int8, 9)
	var a, i, j int8
	a = 0
	for i = 0; i < 3; i++ {
		for j = 0; j < 3; j++ {
			if board[i][j] == "" {
				remaining[a] = []int8{i, j}
				a++
			}
		}
	}

	return remaining, a
}
