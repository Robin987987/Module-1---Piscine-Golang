package piscine

func Sqrt(nb int) int {
	paaritu := 1
	summa := 0
	if nb <= 0 {
		return 0
	} else {
		for ; nb > 0; paaritu = paaritu + 2 {
			nb -= paaritu
			summa++
		}
	}
	if nb == 0 {
		return summa
	} else {
		return 0
	}
}
