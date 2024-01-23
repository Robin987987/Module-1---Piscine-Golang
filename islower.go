package piscine

func IsLower(s string) bool {
	i := 0
	a := false
	for range s {
		if rune(s[i]) >= 97 && rune(s[i]) <= 122 {
			a = true
		} else {
			a = false
			break
		}
		i++
	}
	return a
}
