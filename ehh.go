package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Board represents the Tic-Tac-Toe board
type Board [3][3]string

// Player represents a player's mark (X or O)
type Player string

const (
	PlayerX Player = "U"
	PlayerO Player = "G"
	Empty   string = " "
)

// NewBoard creates and initializes a new empty Tic-Tac-Toe board
func NewBoard() Board {
	var board Board
	for r := 0; r < 3; r++ {
		for c := 0; c < 3; c++ {
			board[r][c] = Empty
		}
	}
	return board
}

// DisplayBoard prints the current state of the board to the console
func (b Board) DisplayBoard() {
	fmt.Println("\n-------------")
	for r := 0; r < 3; r++ {
		fmt.Printf("| %s | %s | %s |\n", b[r][0], b[r][1], b[r][2])
		fmt.Println("-------------")
	}
	fmt.Println()
}

// MakeMove attempts to place a player's mark on the board at the given row and column.
// Returns true if the move was successful, false otherwise (e.g., cell already taken).
func (b *Board) MakeMove(row, col int, player Player) bool {
	if row < 0 || row >= 3 || col < 0 || col >= 3 {
		fmt.Println("Invalid move: Row and column must be between 0 and 2.")
		return false
	}
	if b[row][col] != Empty {
		fmt.Println("Invalid move: That cell is already taken!")
		return false
	}
	b[row][col] = string(player)
	return true
}

// CheckWin checks if the given player has won the game.
func (b Board) CheckWin(player Player) bool {
	mark := string(player)

	// Check rows
	for r := 0; r < 3; r++ {
		if b[r][0] == mark && b[r][1] == mark && b[r][2] == mark {
			return true
		}
	}

	// Check columns
	for c := 0; c < 3; c++ {
		if b[0][c] == mark && b[1][c] == mark && b[2][c] == mark {
			return true
		}
	}

	// Check diagonals
	if (b[0][0] == mark && b[1][1] == mark && b[2][2] == mark) ||
		(b[0][2] == mark && b[1][1] == mark && b[2][0] == mark) {
		return true
	}

	return false
}

// CheckDraw checks if the game is a draw (all cells filled, no winner).
func (b Board) CheckDraw() bool {
	for r := 0; r < 3; r++ {
		for c := 0; c < 3; c++ {
			if b[r][c] == Empty {
				return false // There's an empty cell, so it's not a draw yet
			}
		}
	}
	return true // All cells filled, no winner (checked by CheckWin before calling this)
}

// GetPlayerMove prompts the current player for their move (row and column).
func GetPlayerMove(reader *bufio.Reader, player Player) (int, int, error) {
	fmt.Printf("Player %s, enter your move (row col, e.g., 0 1): ", player)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	parts := strings.Split(input, " ")

	if len(parts) != 2 {
		return 0, 0, fmt.Errorf("invalid input format. Please enter two numbers separated by a space")
	}

	row, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, 0, fmt.Errorf("invalid row: %w", err)
	}

	col, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0, 0, fmt.Errorf("invalid column: %w", err)
	}

	return row, col, nil
}

func main() {
	board := NewBoard()
	currentPlayer := PlayerX
	reader := bufio.NewReader(os.Stdin)

	fmt.("Welcome to Tic-Tac-Toe!")

	for {
		board.DisplayBoard()

		row, col, err := GetPlayerMove(reader, currentPlayer)
		if err != nil {
			fmt.Println(err)
			continue
		}

		if board.MakeMove(row, col, currentPlayer) {
			if board.CheckWin(currentPlayer) {
				board.DisplayBoard()
				fmt.Printf("Player %s wins! Congratulations!\n", currentPlayer)
				break
			}
			if board.CheckDraw() {
				board.DisplayBoard()
				fmt.Println("It's a draw!")
				break
			}
			// Switch players
			if currentPlayer == PlayerX {
				currentPlayer = PlayerO
			} else {
				currentPlayer = PlayerX
			}
		}
	}

	fmt.Println("Thanks for playing!")
}
