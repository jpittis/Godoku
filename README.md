Godoku
===
####A Sudoku solver written in Go.

Though I've written a Sudoku solver before, this one attempts to follow good coding practices as well as proper Go style.

####Status: Complete

####Use Instructions

If you're using OSX, only the executable binary ```godoku``` is required to use this Sudoku solver. All other operating systems must be built with ```go build``` or ```go install```.

The command ```godoku filename.txt``` will attempt to solve the Sudoku puzzle found in ```filename.txt```. If no filename argument is provided, the program will default to ```sudoku.txt```.
