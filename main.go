/* Godoku
A sudoku solver written in Go.

Jake Pittis, December 2014

The goal of this project is to remind me how fun
programming in Go can be! Though I've written a
sudoku solver before, this one attempts to follow
good coding practices as well as proper Go style. */

package main

import (
    "fmt"
    "io/ioutil"
)

/* TODO: This is only for temporary error handling! */
func check(err error) {
    if err != nil {
        panic(err)
    }
}

/* Read in a sudoku file, return converted to an integer array. */
func newSudoku(fileName string) [9][9]int {
    sudoku := [9][9]int{}
    data := readSudoku(fileName)
    parseSudoku(&sudoku, data)
    return sudoku
}

/* Read in a sudoku file.
TODO: handle file errors */
func readSudoku(fileName string) string {
    data, err := ioutil.ReadFile(fileName)
    check(err)
    return string(data)
}

/* Check and place given data into double int array sudoku.
TODO: handle <81 and >81 errors */
func parseSudoku(sudoku *[9][9]int, data string) {
    numbersParsed := 0
    for i := 0; i < len(data); i++ {
        if data[i] >= '0' && data[i] <= '9' && numbersParsed < 81 {
            sudoku[numbersParsed / 9][numbersParsed % 9] = int(data[i] - '0')
            numbersParsed++
        }
    }
}

/* Print a formated sudoku. */
func printSudoku(sudoku *[9][9]int) {
    for i := 0; i < 9; i++ {
        for j := 0; j < 9; j++ {
            fmt.Printf("%d ", sudoku[i][j])
        }
        fmt.Print("\n")
    }
}

/* Return true if given sudoku is correctly solved. */
func checkSudoku(sudoku *[9][9]int) bool {
    for i := 0; i < 9; i++ {
        for j := 0; j < 9; j++ {
            if sudoku[i][j] == 0 || checkAll(sudoku, j, i, sudoku[i][j]) {
                return false
            }
        }
    }
    return true
}

/* Returns true if value placed at x, y is a valid sudoku move. */
func checkAll(sudoku *[9][9]int, x int, y int, value int) bool {
    notInSquare := checkSquare(sudoku, x, y, value)
    notInRow := checkRow(sudoku, x, y, value)
    notInColumn := checkColumn(sudoku, x, y, value)
    return notInSquare && notInRow && notInColumn
}

/* Returns false if value already exists in the 3 by 3 square. */
func checkSquare(sudoku *[9][9]int, x int, y int, value int) bool {
    /* Calculate in which 3 by 3 square x and y are located. */
    x = x - (x % 3)
    y = y - (y % 3)
    /* Check for value inside the square x, y, x + 3, y + 3. */
    for i := y; i < y + 3; i++ {
        for j := x; j < x + 3; j++ {
            if sudoku[i][j] == value {
                return false
            }
        }
    }
    return true
}

/* Returns false if value already exists row. */
func checkRow(sudoku *[9][9]int, x int, y int, value int) bool {
    for i := 0; i < 9; i++ {
        if sudoku[y][i] == value {
            return false
        }
    }
    return true
}

/* Returns false if value slready exists in column. */
func checkColumn(sudoku *[9][9]int, x int, y int, value int) bool {
    for i := 0; i < 9; i++ {
        if sudoku[i][x] == value {
            return false
        }
    }
    return true
}

/* Start solving the given sudoku recursively.
TODO: check sudoku and produce an error if not properly solved*/
func solveSudoku(sudoku *[9][9]int) {
    solveRecursive(sudoku, 0, 0)
}

/* Recursive function for solving sudoku. */
func solveRecursive(sudoku *[9][9]int, x int, y int) bool {
    if x >= 9 || y >= 9 {
        return true
    }

    if sudoku[y][x] != 0 {
        xNext, yNext := nextIndex(x, y)
        return solveRecursive(sudoku, xNext, yNext)
    } else {
        for value := 1; value <= 9; value++ {
            if checkAll(sudoku, x, y, value) {
                sudoku[y][x] = value
                xNext, yNext := nextIndex(x, y)
                if solveRecursive(sudoku, xNext, yNext) {
                    return true
                }
            }
        }
        sudoku[y][x] = 0
        return false
    }
}

/* Increment left to right, up to down. */
func nextIndex(x int, y int) (int, int) {
    if x < 8 {
        x++
    } else {
        x = 0
        y++
    }
    return x, y
}

/* Solves sudoku puzzle found in filename. */
func main() {
    fileName := "sudoku.txt"
    sudoku := newSudoku(fileName)
    fmt.Print("-Given-----------\n")
    printSudoku(&sudoku)
    solveSudoku(&sudoku)
    fmt.Print("-Produced--------\n")
    printSudoku(&sudoku)
    fmt.Print("-----------------\n")
    if checkSudoku(&sudoku) {
        fmt.Print("Correct Solution!\n")
    } else {
        fmt.Print("Something went wrong!\n")
    }
    fmt.Print("-----------------\n")
}