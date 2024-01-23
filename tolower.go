package piscine

func ToLower(s string) string {
	i := 0
	a := []rune(s)
	for range s {
		if rune(s[i]) >= 65 && rune(s[i]) <= 90 {
			a[i] = rune(s[i]) + 32
		}
		i++
	}
	s = string(a)
	return s
}
