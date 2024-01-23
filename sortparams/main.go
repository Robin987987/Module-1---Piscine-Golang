package main

import (
	"os"

	"github.com/01-edu/z01"
)

func SortIntegerTables(table []string) []string {
	i := 1
	for i < len(table) {
		if table[i-1] > table[i] {
			number := table[i]
			table[i] = table[i-1]
			table[i-1] = number
			i = 1
		} else {
			i++
		}
	}
	return table
}

func PrintStrg(s string) {
	vahetatav := []rune(s)
	suurus := len(vahetatav)
	i := 0
	for ; suurus > 0; suurus-- {
		taht := vahetatav[i]
		z01.PrintRune(taht)
		i++
	}
}

func main() {
	bbq := os.Args[1:]
	bbq = SortIntegerTables(bbq)
	for i := 0; i < len(bbq); i++ {
		PrintStrg(bbq[i])
		z01.PrintRune('\n')
	}
}
