package piscine

func NbrBase(nbr int, base string) string {
	number := ""
	if nbr == -9223372036854775808 {
		return "-9223372036854775808"
	} else {
		if len(base) < 2 || !checkUnique(base) || !(Index(base, "+") == -1) || !(Index(base, "-") == -1) {
			return "X"
		} else {
			if nbr < 0 {
				nbr *= -1
				number = Concat(number, "-")
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
				number = Concat(number, string(base[tabel[i]]))
			}
		}
	}
	return number
}
