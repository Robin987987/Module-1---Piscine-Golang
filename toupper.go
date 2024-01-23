package piscine

func ToUpper(s string) string {
	i := 0
	a := []rune(s)
	for range s {
		if rune(s[i]) >= 97 && rune(s[i]) <= 122 {
			a[i] = rune(s[i]) - 32
		}
		i++
	}
	s = string(a)
	return s
}
