package main

import (
	"fmt"
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

func Korrasta(bbq []string) {
	sisesta := 0
	jarg := 0
	var lopp []string
	for a := range bbq[len(bbq)-1] {
		lopp = append(lopp, string(bbq[len(bbq)-1][a]))
	}
	// bbq = SortIntegerTables(bbq)
	for i := 0; i < len(bbq); i++ {
		if len(bbq[i]) > 8 {
			if bbq[i][0:8] == "--insert" {
				sisesta = 1
				for a := range bbq[i][9:] {
					lopp = append(lopp, string(bbq[i][9:][a]))
				}
			}
		}
		if len(bbq[i]) > 6 {
			if bbq[i] == "--order" {
				jarg = 1
			}
		}
		if len(bbq[i]) > 1 {
			if bbq[i][0:2] == "-i" {
				sisesta = 1
				for a := range bbq[i][3:] {
					lopp = append(lopp, string(bbq[i][3:][a]))
				}
			}
		}
		if len(bbq[i]) > 1 {
			if bbq[i] == "-o" {
				jarg = 1
			}
		}
	}

	if sisesta == 1 && jarg == 1 {
		lopp = SortIntegerTables([]string(lopp[0:]))
	} else if sisesta == 0 && jarg == 1 {
		lopp = SortIntegerTables([]string(lopp[0:]))
	} else if sisesta == 1 && jarg == 0 {
		lopp = lopp[0:]
	} else if sisesta == 0 && jarg == 0 {
	}

	for i := range lopp {
		PrintStrg(lopp[i])
	}
	z01.PrintRune('\n')
}

func help() {
	fmt.Println("--insert")
	// z01.PrintRune('\n')
	fmt.Println("  -i")
	// z01.PrintRune('\n')
	fmt.Println("	 This flag inserts the string into the string passed as argument.")
	// z01.PrintRune('\n')
	fmt.Println("--order")
	// z01.PrintRune('\n')
	fmt.Println("  -o")
	// z01.PrintRune('\n')
	fmt.Println("	 This flag will behave like a boolean, if it is called it will order the argument.")
	// z01.PrintRune('\n')
}

func main() {
	bbq := os.Args[1:]
	if len(bbq) < 1 || bbq[0] == "--help" || bbq[0] == "-h" {
		help()
	} else {
		Korrasta(bbq)
	}
}
