package main

import (
	"fmt"
	"os"
)

var board = [9][9]int{}

func showBoard() {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Print(board[i][j])
			if j != 8 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func canPut(x, y, value int) bool {
	return !alreadyInVertical(x, value) &&
		!alreadyInHorizontal(y, value) &&
		!alreadyInSquare(x, y, value)
}

func alreadyInVertical(x, value int) bool {
	for i := range [9]int{} {
		if board[i][x] == value {
			return true
		}
	}
	return false
}

func alreadyInHorizontal(y, value int) bool {
	for i := range [9]int{} {
		if board[y][i] == value {
			return true
		}
	}
	return false
}

func alreadyInSquare(x, y, value int) bool {
	sx, sy := int(x/3)*3, int(y/3)*3
	for dy := range [3]int{} {
		for dx := range [3]int{} {
			if board[sy+dy][sx+dx] == value {
				return true
			}
		}
	}
	return false
}

func next(x, y int) (int, int) {
	nextX, nextY := (x+1)%9, y
	if nextX == 0 {
		nextY = y + 1
	}
	return nextX, nextY
}

func solve(x, y int) bool {
	if y == 9 {
		return true
	}
	if board[y][x] != 0 {
		return solve(next(x, y))
	} else {
		for i := range [9]int{} {
			v := i + 1
			if canPut(x, y, v) {
				board[y][x] = v
				if solve(next(x, y)) {
					return true
				}
				board[y][x] = 0
			}
		}
		return false
	}
}

func checkValues(a [9]int) bool {
	for i := 0; i < 9; i++ {
		for j := i + 1; j < 9; j++ {
			if a[i] > 9 || a[i] < 0 {
				return false
			}
			if a[i] == a[j] && a[i] != 0 && a[j] != 0 {
				return false
			}
		}
	}
	return true
}

func withoutSpace(s string) string {
	res := ""
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			continue
		}
		res = res + string(s[i])
	}
	return res
}

func main() {
	arguments := os.Args[1:]
	if len(arguments) < 1 {
		fmt.Println("Error")
		return
	}
	s := withoutSpace(arguments[0])
	if len(s) == 81 {
		for i := 0; i < len(s); i++ {
			if s[i] == '.' {
				board[i/9][i%9] = 0
			} else {
				board[i/9][i%9] = int(s[i]) - 48
			}
		}
	} else {
		if len(arguments) != 9 {
			fmt.Println("Error")
			return
		} else {
			for i := 0; i < 9; i++ {
				if len(arguments[i]) != 9 {
					fmt.Println("Error")
					return
				}
				for j := 0; j < 9; j++ {
					if arguments[i][j] == '.' {
						board[i][j] = 0
					} else {
						board[i][j] = int(arguments[i][j]) - 48
					}
				}
			}
		}
	}
	for i := 0; i < 9; i++ {
		if !checkValues(board[i]) {
			fmt.Println("Error")
			return
		}
	}
	if solve(0, 0) {
		showBoard()
	} else {
		fmt.Println("Error")
		return
	}
}
