package piscine

func SortRuneTable(table []rune) {
	len := len(table)
	for i := 0; i < len-1; i++ {
		for j := 0; j < len-i-1; j++ {
			if table[j] > table[j+1] {
				table[j], table[j+1] = table[j+1], table[j]
			}
		}
	}
}

func IsValid(base string) bool {
	if len(base) < 2 {
		return false
	}
	runes := []rune(base)
	SortRuneTable(runes)
	for i, rn := range runes {
		if rn == '-' || rn == '+' {
			return false
		}
		if i < len(base)-1 {
			if runes[i] == runes[i+1] { // otsib korduvaid elemente sorteeritud tabelist
				return false
			}
		}
	}
	return true
}

func AtoiBase(s string, base string) int {
	// converts string from given base to
	// base10 int
	var sum int
	if !IsValid(base) {
		return 0
	}
	for i, j := len(s)-1, 0; i >= 0; i-- {
		sum += Index(base, string(s[i])) * RecursivePower(len(base), j)
		// fmt.Printf("index: %v, base: %v string: %v\n", Index(base, string(s[i])), base, string(s[i]))
		j++
	}
	return sum
}
