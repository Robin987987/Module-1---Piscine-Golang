package main

import (
	"os"

	"github.com/01-edu/z01"
)

func PrintStrg(s string) {
	vahetatav := []rune(s)
	suurus := len(vahetatav)
	i := 2
	for ; suurus > 2; suurus-- {
		taht := vahetatav[i]
		z01.PrintRune(taht)
		i++
	}
}

func main() {
	bbq := os.Args
	PrintStrg(bbq[0])
	z01.PrintRune('\n')
}
