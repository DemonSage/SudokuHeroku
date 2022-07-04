package main

import (
	"fmt"
	"math"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"time"
)

var (
	clear map[string]func() //create a map for storing clear funcs
)

func (b Board) displayBoard()  {
	callClear()

	size := len(b.matrix)
	boxEdge := int(math.Sqrt(float64(size)))

	fmt.Println("*Note: The ordinality of rows & columns in the board starts from the top-left corner." +
		"\nRows come first & range from 1 to", size,
		"\nFollowed by columns with the same range of 1 to", size,
		"\nTop-Left corner position: Row 1, Column 1")

	edge := "+"
	for boxes := 0; boxes < boxEdge; boxes++ {
		for slots := 0; slots < boxEdge*2+1; slots++ {
			edge += "-"
		}
		edge += "+"
	}

	for rows := 0; rows <= size; rows++ {
		if rows % boxEdge == 0 {
			fmt.Println(edge)
		}
		if rows != size {
			for columns := 0; columns <= size; columns++ {
				if columns % boxEdge == 0 {
					fmt.Print("| ")
				}
				if columns != size {
					fmt.Printf("%v ",b.matrix[rows][columns])
				}
			}
			fmt.Println()
		}
	}
}

func init() {
	clear = make(map[string]func()) //Initialize it
	clear["linux"] = func() {
		cmd := exec.Command("clear") //Linux createFile, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") //Windows createFile, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

//clears terminal screen, uses init()
func callClear() {
	value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	if ok { //if we defined a clear func for that platform:
		value()  //we execute it
	} else { //unsupported platform
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}

func loadingBar() {
	time.Sleep(450 * time.Millisecond)
	fmt.Print(".")
	time.Sleep(450 * time.Millisecond)
	fmt.Print(".")
	time.Sleep(450 * time.Millisecond)
	fmt.Print(".")
	time.Sleep(700 * time.Millisecond)
	fmt.Println()
	callClear()
}

//the format for using this should preclude/exclude any Println or "\n" before usage
func optionsMenu(options ...string) string {
	//display menu portion
	fmt.Println("\n————————————————————")
	for i, s := range options {
		num := i+1
		fmt.Printf("(%v) | %v\n", num, s)
	}
	fmt.Println("(-)  Exit Game" +
		"\n————————————————————")

	//user input portion
	userChoice := scan()
	if userChoice == "-" { //makes exit option the first thing the program checks
		callClear()
		fmt.Println("Goodbye!")
		loadingBar()
		os.Exit(0)
	}

	valid := false
	for !valid {
		for i, _ := range options {
			validChoice := strconv.Itoa(i+1) //increments index by 1 to be consistent with ui, then converts to string to be compared
			if userChoice == validChoice {
				valid = true //breaks out of loop
			}
		}
		if !valid {
			fmt.Println("Invalid input. Please try again.") //default case if user input is not a valid option
			userChoice = scan()
		}
	}
	return userChoice
}