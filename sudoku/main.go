package main

import (
	"fmt"
	"os"
)

func main() {
	arr := os.Args[1:]
	if checkErrors(arr) {
		fmt.Println("Error")
	} else {
		sudoku := optimizeSudoku(arr)
		if solveSudoku(&sudoku, len(sudoku)) {
			printSudoku(sudoku)
		} else {
			fmt.Println("Error")
		}
	}
}

func checkErrors(arr []string) bool { // funksioon errorite kontrollimiseks
	for i, str := range arr {
		if len(str) != 9 {
			return true
		}
		for _, ch := range arr[i] {
			if !(ch >= '1' && ch <= '9' || ch == '.') {
				return true
			}
		}
	}
	return false
}

func solveSudoku(arr *[][]int, length int) bool { // main funksioon sudoku lahendamiseks
	isEmpty := true
	row := -1
	column := -1
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			if (*arr)[i][j] == 0 {
				row = i
				column = j
				isEmpty = false
				break
			}
		}
	}
	if isEmpty {
		return true
	}
	for number := 1; number <= 9; number++ {
		if isCorrect(*arr, row, column, number) {
			(*arr)[row][column] = number

			if solveSudoku(arr, length) {
				return true
			} else {
				(*arr)[row][column] = 0
			}
		}
	}
	return false
}

func optimizeSudoku(arr []string) [][]int { // veidike sudoku optimiseerimist
	sudoku := make([][]int, 9)
	for i := range sudoku {
		sudoku[i] = make([]int, 9)
	}
	for i, str := range arr {
		for j, ch := range str {
			sudoku[i][j] = runeToInt(ch)
		}
	}
	return sudoku
}

func printSudoku(arr [][]int) { // sudoku printimine commandline-le
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			fmt.Print(arr[i][j])
			fmt.Print(" ")
		}
		fmt.Println()
	}
}

func isCorrect(arr [][]int, row int, column int, num int) bool { // fuksioon sudoku kontrollimiseks
	return !checkRow(arr, row, num) && !checkColumn(arr, column, num) && !checkSubSudoku(arr, row-(row%3), column-(column%3), num)
}

func checkRow(arr [][]int, row int, num int) bool { // funksioon rea kontrollimiseks
	for column := 0; column < len(arr); column++ {
		if arr[row][column] == num {
			return true
		}
	}
	return false
}

func checkColumn(arr [][]int, column int, num int) bool { // unksioon tulba kontrollimiseks
	for row := 0; row < len(arr); row++ {
		if arr[row][column] == num {
			return true
		}
	}
	return false
}

func checkSubSudoku(arr [][]int, rowIndex int, columnIndex int, num int) bool { // sub sudoku kontrollimine
	for row := 0; row < 3; row++ {
		for column := 0; column < 3; column++ {
			if arr[rowIndex+row][columnIndex+column] == num {
				return true
			}
		}
	}
	return false
}

func runeToInt(number rune) int { // muudab rune int-ks
	count := 0
	for i := '1'; i <= number; i++ {
		count++
	}
	return count
}
