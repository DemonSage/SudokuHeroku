package main

func debug() {

	//displayBoard()

	//var testRow0 = []int{0,0,0,0,0,0,0,0,0}
	//a, b := checkRow(testRow0)
	//fmt.Println("Status:", a, "Code:", b)

	var (
	/*
		testBoardIncomplete = Board{
			[][]int{
				{4,0,0, 0,0,9, 0,0,0},
				{1,2,3, 4,5,6, 7,8,9},
				{4,0,0, 0,0,3, 0,0,0},

				{4,5,5, 1,3,2, 0,0,0},
				{7,8,7, 6,5,4, 3,2,1},
				{9,5,5, 9,7,8, 0,0,0},

				{8,0,0, 0,0,7, 0,0,0},
				{2,0,0, 0,0,5, 0,0,0},
				{3,0,0, 0,0,1, 0,0,0},
			},
		}
		testBoardFilledIncorrect = Board{
			[][]int{
				{1,2,3, 4,5,6, 7,8,9},
				{1,2,3, 4,5,6, 7,8,9},
				{1,2,3, 4,5,6, 7,8,9},

				{1,2,3, 4,5,6, 7,8,9},
				{1,2,3, 4,5,6, 7,8,9},
				{1,2,3, 4,5,6, 7,8,9},

				{1,2,3, 4,5,6, 7,8,9},
				{1,2,3, 4,5,6, 7,8,9},
				{1,2,3, 4,5,6, 7,8,9},
			},
		}
		testBoardNormalCorrect = Board{
			[][]int{
				{5,8,1, 6,7,2, 4,3,9},
				{7,9,2, 8,4,3, 6,5,1},
				{3,6,4, 5,9,1, 7,8,2},

				{4,3,8, 9,5,7, 2,1,6},
				{2,5,6, 1,8,4, 9,7,3},
				{1,7,9, 3,2,6, 8,4,5},

				{8,4,5, 2,1,9, 3,6,7},
				{9,1,3, 7,6,8, 5,2,4},
				{6,2,7, 4,3,5, 1,9,8},
			},
		}
	*/

	/*testBoardEasyIncomplete = Board{
		[][]int{
			{4,0,0,0},
			{1,2,3,4},
			{4,0,0,0},
			{4,5,5,1},
		},
		nil,
	}
	testBoardEasyIncorrect = Board{
		[][]int{
			{4,3,2,1},
			{1,2,3,4},
			{4,3,2,3},
			{4,5,5,1},
		},
		nil,
	}
	testBoardEasyIncorrect2 = Board{
		[][]int{
			{1,2,3,4},
			{2,3,4,1},
			{3,4,1,2},
			{4,3,2,1},
		},
		nil,
	}*/

	/*testBoardEasyCorrect = Board{
		[][]int{
			{1,4,3,2},
			{3,2,1,4},
			{4,1,2,3},
			{2,3,4,1},
		},
		nil,
	}*/

	/*testBoardNormalCorrect = Board{
		[][]int{
			{5, 8, 1, 6, 7, 2, 4, 3, 9},
			{7, 9, 2, 8, 4, 3, 6, 5, 1},
			{3, 6, 4, 5, 9, 1, 7, 8, 2},

			{4, 3, 8, 9, 5, 7, 2, 1, 6},
			{2, 5, 6, 1, 8, 4, 9, 7, 3},
			{1, 7, 9, 3, 2, 6, 8, 4, 5},

			{8, 4, 5, 2, 1, 9, 3, 6, 7},
			{9, 1, 3, 7, 6, 8, 5, 2, 4},
			{6, 2, 7, 4, 3, 5, 1, 9, 8},
		},
		nil,
	}*/

	/*testBoardHardCorrect = Board{
		[][]int{
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},

			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},

			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},

			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		nil,
	}*/
	)

	/*
		a, b := checkColumn(testBoardIncomplete.getColumn(6))
		fmt.Println("Status:", a, "Code:", b)
		a, b = checkColumn(testBoardIncomplete.getColumn(1))
		fmt.Println("Status:", a, "Code:", b)
		a, b = checkColumn(testBoardIncomplete.getColumn(9))
		fmt.Println("Status:", a, "Code:", b)

		a, b := testBoardIncomplete.checkBoard()
		fmt.Println("Status:", a, "Code:", b)
		a, b = testBoardFilledIncorrect.checkBoard()
		fmt.Println("Status:", a, "Code:", b)
		a, b = testBoardNormalCorrect.checkBoard()
		fmt.Println("Status:", a, "Code:", b)
	*/

	/*testBoardEasyIncomplete.getOrigins()
	testBoardEasyIncorrect.getOrigins()
	testBoardEasyIncorrect2.getOrigins()*/

	//4x4 board tests
	//should be 0, 1, 1, -1 respectively
	/*a, b := testBoardEasyIncomplete.checkBoard()
	fmt.Println("Status:", a, "Code:", b)
	a, b = testBoardEasyIncorrect.checkBoard()
	fmt.Println("Status:", a, "Code:", b)
	a, b = testBoardEasyIncorrect2.checkBoard()
	fmt.Println("Status:", a, "Code:", b)*/

	/*//test 4x4
	testBoardEasyCorrect.displayBoard()
	testBoardEasyCorrect.getOrigins()
	a, b := testBoardEasyCorrect.checkBoard()
	fmt.Println("Status:", a, "Code:", b)

	//test 9x9
	testBoardNormalCorrect.displayBoard()
	testBoardNormalCorrect.getOrigins()
	a, b = testBoardNormalCorrect.checkBoard()
	fmt.Println("Status:", a, "Code:", b)

	//test 9x9
	testBoardHardCorrect.displayBoard()
	testBoardHardCorrect.getOrigins()
	a, b = testBoardHardCorrect.checkBoard()
	fmt.Println("Status:", a, "Code:", b)*/
}
