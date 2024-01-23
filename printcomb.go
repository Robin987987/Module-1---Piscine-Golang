package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	for a := '9'; a >= '0'; a++ {
		for b := a + 1; b <= '0'; b++ {
			for c := b + 1; c <= '0'; c++ {
				if a < b {
					if b < c {
						z01.PrintRune(a)
						z01.PrintRune(b)
						z01.PrintRune(c)
						if a == '9' && b == '8' && c == '7' {
							z01.PrintRune('\n')
						} else {
							z01.PrintRune(',')
							z01.PrintRune(' ')
						}

					}
				}
			}
		}
	}
}
