package piscine

func IsPrintable(s string) bool {
	i := 0
	a := false
	for range s {
		if rune(s[i]) >= 32 && rune(s[i]) <= 127 {
			a = true
		} else {
			a = false
			break
		}
		i++
	}
	return a
}
