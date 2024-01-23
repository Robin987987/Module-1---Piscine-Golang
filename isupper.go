package piscine

func IsUpper(s string) bool {
	i := 0
	a := false
	for range s {
		if rune(s[i]) >= 65 && rune(s[i]) <= 90 {
			a = true
		} else {
			a = false
			break
		}
		i++
	}
	return a
}
