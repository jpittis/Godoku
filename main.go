/* Godoku
A sudoku solver written in Go.

Jake Pittis, December 2014

Though I've written a sudoku solver before, this
one attempts to follow good coding practices as
well as proper Go style. */

package main

import (
    "fmt"
    "io/ioutil"
    "errors"
    "os"
)

/* Read in a sudoku file, return converted to an integer array. */
func newSudoku(fileName string) ([9][9]int, error) {
    sudoku := [9][9]int{}

    data, err := readSudoku(fileName)
    if err != nil {
        return sudoku, err
    }

    err = parseSudoku(&sudoku, data)
    if err != nil {
        return sudoku, err
    }
    return sudoku, nil
}

/* Read in a sudoku file. */
func readSudoku(fileName string) (string, error) {
    data, err := ioutil.ReadFile(fileName)
    if err != nil {
        return string(data), fmt.Errorf("Error when reading file '%s' containing sudoku.", fileName)
    } else {
        return string(data), nil
    }
}

/* Check and place given data into double int array sudoku. */
func parseSudoku(sudoku *[9][9]int, data string) error {
    numbersParsed := 0
    for i := 0; i < len(data); i++ {
        if data[i] >= '0' && data[i] <= '9' && numbersParsed < 81 {
            sudoku[numbersParsed / 9][numbersParsed % 9] = int(data[i] - '0')
            numbersParsed++
        }
    }
    if numbersParsed != 81 {
        return errors.New("Input file must have exactly 81 numbers to be a valid sudoku puzzle.")
    }
    return nil
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
    return checkSquare(sudoku, x, y, value) &&
    checkRow(sudoku, x, y, value) &&
    checkColumn(sudoku, x, y, value)
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

/* Start solving the given sudoku recursively. */
func solveSudoku(sudoku *[9][9]int) error {
    solveRecursive(sudoku, 0, 0)
    if checkSudoku(sudoku) {
        return nil
    } else {
        return errors.New("Attempt to solve sudoku has failed. Maybe the puzzle inputed has no solution!")
    }
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
    args := os.Args
    var fileName string

    if len(args) <= 1 {
        fileName = "sudoku.txt"
    } else {
        fileName = args[1]
    }

    sudoku, err := newSudoku(fileName)
    if err != nil {
        fmt.Print("Error: ", err.Error(), "\n")
        return
    }

    err = solveSudoku(&sudoku)
    if err != nil {
        fmt.Print("Error: ", err.Error(), "\n")
        return
    }

    printSudoku(&sudoku)
}