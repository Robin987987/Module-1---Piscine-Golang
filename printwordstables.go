package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
	for i := range a {
		for j := range a[i] {
			z01.PrintRune(rune(a[i][j]))
		}
		z01.PrintRune('\n')
	}
}
