package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\t')
	if name == "" {
		fmt.Println("Not a Raid function")
		return
	}
	columns := 0
	rows := 0
	for _, v := range name {
		if v == '\n' {
			rows++
		} else {
			columns++
		}
	}
	columns = columns / rows
	first := true
	raid := false
	if name == quadA(rows, columns) {
		fmt.Print("[quadA] [", columns, "] [", rows, "]")
		first = false
		raid = true
	}
	if name == quadB(rows, columns) {
		if !first {
			fmt.Print(" || ")
		}
		fmt.Print("[quadB] [", columns, "] [", rows, "]")
		first = false
		raid = true
	}
	if name == quadC(rows, columns) {
		if !first {
			fmt.Print(" || ")
		}
		fmt.Print("[quadC] [", columns, "] [", rows, "]")
		first = false
		raid = true
	}
	if name == quadD(rows, columns) {
		if !first {
			fmt.Print(" || ")
		}
		fmt.Print("[quadD] [", columns, "] [", rows, "]")
		first = false
		raid = true
	}
	if name == quadE(rows, columns) {
		if !first {
			fmt.Print(" || ")
		}
		fmt.Print("[quadE] [", columns, "] [", rows, "]")
		first = false
		raid = true
	}
	if !raid {
		fmt.Println("Not a Raid function")
	} else {
		fmt.Println()
	}
}

func quadA(h, v int) string {
	result := ""
	for i := 1; i <= h; i++ {
		for j := 1; j <= v; j++ {
			if i == 1 && j == 1 || i == 1 && j == v || i == h && j == v || i == h && j == 1 {
				result += "o"
			} else if i == 1 || i == h {
				result += "-"
			} else if j == 1 || j == v {
				result += "|"
			} else {
				result += " "
			}
		}
		result += "\n"
	}
	return result
}

func quadB(h, v int) string {
	result := ""
	for i := 1; i <= h; i++ {
		for j := 1; j <= v; j++ {
			if i == 1 && j == 1 {
				result += "/"
			} else if i == 1 && j == v || i == h && j == 1 {
				result += "\\"
			} else if i == h && j == v {
				result += "/"
			} else if i == 1 || i == h || j == 1 || j == v {
				result += "*"
			} else {
				result += " "
			}
		}
		result += "\n"
	}
	return result
}

func quadC(h, v int) string {
	result := ""
	for i := 1; i <= h; i++ {
		for j := 1; j <= v; j++ {
			if i == 1 && (j == 1 || j == v) {
				result += "A"
			} else if i == h && (j == 1 || j == v) {
				result += "C"
			} else if i == 1 || i == h || j == 1 || j == v {
				result += "B"
			} else {
				result += " "
			}
		}
		result += "\n"
	}
	return result
}

func quadD(h, v int) string {
	result := ""
	for i := 1; i <= h; i++ {
		for j := 1; j <= v; j++ {
			if j == 1 && (i == 1 || i == h) {
				result += "A"
			} else if j == v && (i == 1 || i == h) {
				result += "C"
			} else if i == 1 || i == h || j == 1 || j == v {
				result += "B"
			} else {
				result += " "
			}
		}
		result += "\n"
	}
	return result
}

func quadE(h, v int) string {
	result := ""
	for i := 1; i <= h; i++ {
		for j := 1; j <= v; j++ {
			if j == 1 && i == 1 {
				result += "A"
			} else if (j == v && i == 1) || (i == h && j == 1) {
				result += "C"
			} else if j == v && i == h {
				result += "A"
			} else if i == 1 || i == h || j == 1 || j == v {
				result += "B"
			} else {
				result += " "
			}
		}
		result += "\n"
	}
	return result
}
