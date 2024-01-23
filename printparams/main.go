package main

import (
	"os"

	"github.com/01-edu/z01"
)

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
	bbq := os.Args
	for i := 1; i < len(bbq); i++ {
		PrintStrg(bbq[i])
		z01.PrintRune('\n')
	}
}
