package piscine

func IsNumeric(s string) bool {
	i := 0
	a := false
	for range s {
		if rune(s[i]) >= 48 && rune(s[i]) <= 57 {
			a = true
		} else {
			a = false
			break
		}
		i++
	}
	return a
}
