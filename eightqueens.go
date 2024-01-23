package piscine

import "github.com/01-edu/z01"

var (
	vastus  [8]rune
	tõde    [9]bool
	vastus2 [9]int
)

func EightQueens() {
	ok := true
	cnt := 0
	for i := 1; i <= 8; i++ {
		if tõde[i] == false {
			ok = false
		} else {
			cnt++
		}
	}

	if ok == true {
		for _, c := range vastus {
			z01.PrintRune(c)
		}
		z01.PrintRune('\n')

		return
	}
	for i := '1'; i <= '8'; i++ {
		cur := 0
		for j := '1'; j <= i; j++ {
			cur++
		}
		if tõde[cur] == false {
			put := true
			for j := 1; j <= cnt; j++ {
				if cur == vastus2[j]-(cnt+1-j) || cur == vastus2[j]+(cnt+1-j) {
					put = false
					break
				}
			}
			if put == true {
				tõde[cur] = true
				vastus[cnt] = i
				vastus2[cnt+1] = cur
				EightQueens()
				tõde[cur] = false

			}
		}
	}
}
