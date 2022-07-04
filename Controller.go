package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	gameBoard = Board{
		matrix:  nil,
		origins: nil,
	}
)

func main() {
	startUp()
	playGame()
}


func buildSudokuBoard() Board {
	b := Board{
		matrix:  nil,
		origins: nil,
	}

	//Easy = 4x4
	//Normal = 9x9
	//Hard = 16x16
	//Extreme = 25x25
	//Sudoku-zilla = 100x100
	fmt.Println("Choose a difficulty: ")
	choice := optionsMenu("Easy", "Normal", "Hard", "Extreme", "Sudoku-zilla")
	switch choice {
	case "1":
		//4x4
		b.matrix = make([][]int, 4)
		for i := range b.matrix {
			b.matrix[i] = make([]int, 4)
		}
		b.getOrigins()
	case "2":
		//9x9
		b.matrix = make([][]int, 9)
		for i := range b.matrix {
			b.matrix[i] = make([]int, 9)
		}
		b.getOrigins()
	case "3":
		//16x16
		b.matrix = make([][]int, 16)
		for i := range b.matrix {
			b.matrix[i] = make([]int, 16)
		}
		b.getOrigins()
	case "4":
		//25x25
		b.matrix = make([][]int, 25)
		for i := range b.matrix {
			b.matrix[i] = make([]int, 25)
		}
		b.getOrigins()
	case "5":
		callClear()
		fmt.Println("wip")
		os.Exit(343)
	}
	return b
}

//Utility func to make getting user input simple
func scan() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func startUp() {
	callClear()
	fmt.Println("Welcome to SudokuHeroku!\n\"The best sudoku app on this side of the galaxy\"")
	fmt.Println("\nWhat would you like to do?")

	choice := optionsMenu("Play", "Solve my board")
	switch choice {
	case "1":
		callClear()
		gameBoard = buildSudokuBoard()
		gameBoard.displayBoard()
	case "2":
		callClear()
		fmt.Println("this feature is coming soon")
	}
}

func playGame() {
	maxValidValue := len(gameBoard.matrix)
	gameWon := false
	for !gameWon {
		gameOptions(maxValidValue)
		gameBoard.displayBoard()
		gameWon, _ = gameBoard.checkBoard()
	}

	callClear()
	fmt.Println("Congratz, You won!")
}

func gameOptions(m int) {
	fmt.Println("Please enter the desired row:")
	chosenRow := scan()
	for !validNumber(chosenRow, m) {
		fmt.Println("Please enter a valid number.")
		chosenRow = scan()
	}

	fmt.Println("Please enter the desired column:")
	chosenColumn := scan()
	for !validNumber(chosenColumn, m) {
		fmt.Println("Please enter a valid number.")
		chosenColumn = scan()
	}

	fmt.Printf("Please enter a number between 1-%v:\n", m)
	chosenNum := scan()
	for !validNumber(chosenNum, m) {
		fmt.Println("Please enter a valid number.")
		chosenNum = scan()
	}

	row, _ := strconv.Atoi(chosenRow)
	column, _ := strconv.Atoi(chosenColumn)
	num, _ := strconv.Atoi(chosenNum)

	gameBoard.changeBoard(row-1, column-1, num)
}

func validNumber(num string, m int) bool {
	for i := 1; i <= m; i++ {
		validChoice := strconv.Itoa(i) //converts to string to be compared
		if num == validChoice {
			return true
		}
	}
	return false
}
