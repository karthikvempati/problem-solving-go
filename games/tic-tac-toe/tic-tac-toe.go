package tic_tac_toe

import (
	"fmt"
	"math/rand"
)

func IsGameOver(board [][]int8) (bool, int8) {
	for i := 0; i < 3; i++ {
		if board[i][0] == board[i][1] && board[i][0] == board[i][2] && board[i][0] != -1 {
			return true, board[i][0]
		}
	}
	for i := 0; i < 3; i++ {
		if board[0][i] == board[1][i] && board[0][i] == board[2][i] && board[0][i] != -1 {
			return true, board[0][i]
		}
	}
	if board[0][0] == board[1][1] && board[0][0] == board[2][2] && board[0][0] != -1 {
		return true, board[0][0]
	}

	if board[0][2] == board[1][1] && board[0][2] == board[2][0] && board[1][1] != -1 {
		return true, board[0][2]
	}

	return false, -1
}

func MakeMove(i, j int8, board [][]int8, player int8) ([][]int8, bool) {
	moved := false
	if i < 0 || j < 0 || i > 2 || j > 2 {
		fmt.Println("Please enter a valid move")
	} else if board[i][j] != -1 {
		fmt.Println("Please enter a valid move")
	} else {
		board[i][j] = player
		moved = true
		fmt.Println("Player ", player, " move", i, j)
		PrintBoard(board)
	}

	return board, moved
}

func PrintBoard(board [][]int8) {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] != -1 {
				fmt.Print(" ")
			}
			fmt.Print(board[i][j], " ")
		}
		fmt.Println()
	}
}

func InitializeGame() {
	board := make([][]int8, 3)
	for i := range board {
		board[i] = []int8{-1, -1, -1}
	}
	// board := [][]int8{{-1, -1, -1}, {-1, -1, -1}, {-1, -1, -1}}
	fmt.Print("Enter text: ")
	var i, j int8
	//fmt.Scanf("%d %d", &i, &j)
	var human, computer int8 = 0, 1
	var moved bool
	prevPlayer := int8(computer)
	for a := 0; a < 9; {
		gameOver, player := IsGameOver(board)
		if gameOver {
			fmt.Println("Game over, Player", player, "wins")
			break
		}
		if prevPlayer == int8(human) {
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

func ChangePlayer(prevPlayer int8) int8 {
	if prevPlayer == 0 {
		return 1
	} else if prevPlayer == 1 {
		return 0
	} else {
		panic("Not a valid player")
	}
}

func RandomCoordinates(board [][]int8) (int8, int8) {
	a, b := GetRemainingSlots(board)
	randVal := rand.Intn(int(b))
	return a[randVal][0], a[randVal][1]
}

func GetRemainingSlots(board [][]int8) ([][]int8, int8) {
	remaining := make([][]int8, 9)
	var a, i, j int8
	a = 0
	for i = 0; i < 3; i++ {
		for j = 0; j < 3; j++ {
			if board[i][j] == -1 {
				remaining[a] = []int8{i, j}
				a++
			}
		}
	}

	return remaining, a
}
