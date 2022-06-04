package main

import (
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func printStr(str string) {
	r := []rune(str)

	for i := 0; i < len(r); i++ {
		z01.PrintRune(r[i])
	}
}

func main() {
	if len(os.Args) == 1 {
		bytes, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			return
		}
		printStr(string(bytes))
	}
	for i := 1; i < len(os.Args); i++ {
		data, err := ioutil.ReadFile(os.Args[i])
		if err != nil {
			printStr("ERROR: open " + os.Args[i] + ": no such file or directory" + "\n")
			os.Exit(1)
			break
		}
		printStr(string(data))
	}
}
