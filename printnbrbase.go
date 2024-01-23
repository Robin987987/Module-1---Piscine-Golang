package piscine

import "github.com/01-edu/z01"

func checkUnique(s string) bool {
	for i := 0; i < len(s)-1; i++ {
		for j := i + 1; j < len(s); j++ {
			if rune(s[i]) == rune(s[j]) {
				return false
			}
		}
	}
	return true
}

func PrintNbrBase(nbr int, base string) {
	if nbr == -9223372036854775808 {
		z01.PrintRune('-')
		z01.PrintRune('9')
		z01.PrintRune('2')
		z01.PrintRune('2')
		z01.PrintRune('3')
		z01.PrintRune('3')
		z01.PrintRune('7')
		z01.PrintRune('2')
		z01.PrintRune('0')
		z01.PrintRune('3')
		z01.PrintRune('6')
		z01.PrintRune('8')
		z01.PrintRune('5')
		z01.PrintRune('4')
		z01.PrintRune('7')
		z01.PrintRune('7')
		z01.PrintRune('5')
		z01.PrintRune('8')
		z01.PrintRune('0')
		z01.PrintRune('8')
	} else {
		if len(base) < 2 || !checkUnique(base) || !(Index(base, "+") == -1) || !(Index(base, "-") == -1) {
			z01.PrintRune('N')
			z01.PrintRune('V')
		} else {
			if nbr < 0 {
				nbr *= -1
				z01.PrintRune('-')
			}
			var tabel []rune
			uus := nbr
			pikkus := len(base)
			for {
				jaak := uus % pikkus
				uus /= pikkus
				tabel = append(tabel, rune(jaak))
				if uus == 0 {
					break
				}
			}
			for i := len(tabel) - 1; i >= 0; i-- {
				z01.PrintRune(rune(base[tabel[i]]))
			}
		}
	}
}
