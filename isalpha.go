package piscine

func IsAlpha(s string) bool {
	i := 0
	a := false
	for range s {
		if rune(s[i]) >= 65 && rune(s[i]) <= 90 || rune(s[i]) >= 97 && rune(s[i]) <= 122 || rune(s[i]) >= 48 && rune(s[i]) <= 57 {
			a = true
		} else {
			a = false
			break
		}
		i++
	}
	return a
}
