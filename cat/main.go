package main

import (
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		data, _ := ioutil.ReadAll(os.Stdin)
		PrintStuff(string(data))
	}
	EOF := false
	for _, s := range arguments[1:] {
		if s == "<<EOF>" {
			EOF = true
		} else if s == "cat" || s == "./cat" {
			// Meow
		} else {
			textFile, err := ioutil.ReadFile(s)
			if err != nil {
				printMessgage := "ERROR: open " + s + ": no such file or directory"
				PrintStuff(printMessgage)
				z01.PrintRune('\n')
				os.Exit(1)
			} else {
				PrintStuff(string(textFile))
			}
		}
	}
	if EOF {
		z01.PrintRune('\n')
		PrintStuff("EOF")
	}
}

func PrintStuff(s string) {
	for _, character := range s {
		z01.PrintRune(character)
	}
}
