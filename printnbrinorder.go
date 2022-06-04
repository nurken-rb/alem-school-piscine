package piscine

import (
	"github.com/01-edu/z01"
)

func defSize(n int) int {
	size := 0

	for n != 0 {
		n /= 10
		size++
	}
	return size
}

func makeDigitArr(n int) []int {
	arr := make([]int, defSize(n))

	for i := 0; n != 0; i++ {
		arr[i] = n % 10
		n /= 10
	}
	return arr
}

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune(48)
	}
	arr := makeDigitArr(n)

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				temp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
			}
		}
	}

	for i := 0; i < len(arr); i++ {
		z01.PrintRune(rune(arr[i]) + 48)
	}
}
