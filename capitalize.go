package piscine

func Capitalize(s string) string {
	i := 0
	a := []rune(s)
	sona := 0
	for range s {
		if (rune(s[i]) >= 65 && rune(s[i]) <= 90 || rune(s[i]) >= 48 && rune(s[i]) <= 57) && sona == 0 {
			sona = 1
		} else if rune(s[i]) >= 97 && rune(s[i]) <= 122 && sona == 0 {
			a[i] = rune(s[i]) - 32
			sona = 1
		} else if rune(s[i]) >= 65 && rune(s[i]) <= 90 {
			a[i] = rune(s[i]) + 32
		} else if !(IsAlpha(string(s[i]))) {
			sona = 0
		}
		i++
	}
	s = string(a)
	return s
}
