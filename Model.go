package main

import (
	"math"
)

type point struct {
	x, y int
}

type Board struct {
	matrix [][]int
	origins []point
}

//b *Board is a pointer receiver which points to the original struct
//whereas b Board would simply be a value receiver, working on a copy of the original struct
func (b *Board) getOrigins() {
	size := len(b.matrix)
	originSlice := make([]point, size+1)
	originSlice[0] = point{-1,-1} //sentinel point for re-indexing to match human abstraction

	//need the constant increment value
	//by which to discover origin points for boxes
	incrementFloat := float64(size)
	incrementFloat = math.Sqrt(incrementFloat)
	increment := int(incrementFloat)

	//Works for 4x4 & 9x9
	index := 1
	for i := 0; i < increment; i++ {
		for j := 0; j < increment; j++ {
			originSlice[index] = point{i*increment,j*increment}
			index++
		}
	}

	b.origins = originSlice
}

func (b *Board) changeBoard(row int, column int, value int) {
	b.matrix[row][column] = value
}

/*
Will take a row slice from a Board struct
and check the validity of the row
by returning true if valid, false otherwise
and returning a code
on why it is false or true
	-1:	true + valid
	0: 	false + contains an empty box
	1: 	false + contains incorrect values
*/
func checkRow(row []int) (bool, int) {
	//make a variable length slice to store statuses of the row
	//default initialization state should be all false for boolSlice
	var boolSlice = make([]bool, len(row))
	for x := 0; x < len(row); x++ {
		if row[x] == 0 {
			return false, 0
		}
		/*
			find whatever value row[x] contains.
			that value is the index we want for boolSlice.
			go to the corresponding index in boolSlice
			then flip the status of that bool
		*/
		boolSlice[row[x]-1] = !boolSlice[row[x]-1]
	}
	for _, s := range boolSlice {
		if s == false {
			return false, 1
		}
	}
	return true, -1 //passed all disqualifying conditions at this point
}

func (b  Board) getColumn(column int) []int {
	var constructedColumn = make([]int, len(b.matrix))
	for x := 0; x < len(b.matrix); x++ {
		constructedColumn[x] = b.matrix[x][column-1]
	}
	return constructedColumn
}

//logic identical to checkRow
func checkColumn(column []int) (bool, int) {
	var boolSlice = make([]bool, len(column))
	for x := 0; x < len(column); x++ {
		if column[x] == 0 {
			return false, 0
		}
		/*
			find whatever value column[x] contains.
			that value is the index we want for boolSlice.
			go to the corresponding index in boolSlice
			then flip the status of that bool
		*/
		boolSlice[column[x]-1] = !boolSlice[column[x]-1]
	}
	for _, b := range boolSlice {
		if b == false {
			return false, 1
		}
	}
	return true, -1
}

//uses point slice origins
func (b  Board) getBox(box int) [][]int {

	boxSizeInt := len(b.matrix)
	boxSizeFloat := float64(boxSizeInt)
	boxSizeFloat = math.Sqrt(boxSizeFloat)
	boxSize := int(boxSizeFloat)

	//refactor this using slice logic
	i := b.origins[box].x
	j := b.origins[box].y

	constructedBox := make([][]int, boxSize)

	for rows := 0; rows < boxSize; rows++ {
		//slicing the right sections
		constructedBox[rows] = b.matrix[i][j:j+boxSize]
		i++ //moves the index to the next row in the original sudoku Board
	}
	return constructedBox
}

func checkBox(box [][]int) (bool, int) {
	var boolSlice = make([]bool, len(box)*len(box))
	for _, s := range box {
		for _, t := range s {
			if t == 0 {
				return false, 0
			}
			boolSlice[t-1] = !boolSlice[t-1]
		}
	}
	for _, s := range boolSlice {
		if s == false {
			return false, 1
		}
	}
	return true, -1
}

func (b Board) checkBoard() (bool, int) {
	//check all rows
	for i := 0; i < len(b.matrix); i++ {
		a, b := checkRow(b.matrix[i])
		if a == false {
			return a, b
		}
	}

	//check all columns
	for i := 1; i <= len(b.matrix); i++ {
		a, b := checkColumn(b.getColumn(i))
		if a == false {
			return a, b
		}
	}

	//check all boxes
	for i := 1; i <= len(b.matrix); i++ {
		a, b := checkBox(b.getBox(i))
		if a == false {
			return a, b
		}
	}

	return true, -1 //The sudoku board has been Solved  correctly
}