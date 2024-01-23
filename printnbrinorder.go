package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbrInOrder(n int) {
	if n < 10 {
		z01.PrintRune(rune(n + 48))
	} else {
		var tabel []rune
		kordus := 1
		arv := 0
		s := n
		o := 0
		j := '0'
		i := 0
		y := 1
		for s != 0 {
			arv++
			s /= 10
		}

		for ; arv > 1; arv-- {
			kordus = kordus * 10
		}
		for ; kordus >= 1; kordus = kordus / 10 {
			k := n / kordus
			for o = k; o >= 1; o-- {
				j++
			}
			tabel = append(tabel, j)
			i++
			j = '0'
			o := 0
			o = k * kordus
			n = n - o

		}
		for y < len(tabel) {
			if tabel[y-1] > tabel[y] {
				number := tabel[y]
				tabel[y] = tabel[y-1]
				tabel[y-1] = number
				y = 1
			} else {
				y++
			}
		}
		for b := range tabel {
			z01.PrintRune(tabel[b])
		}
	}
}
