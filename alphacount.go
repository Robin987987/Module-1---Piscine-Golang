package piscine

func AlphaCount(s string) int {
	i := 0
	pikkus := 0
	for ; i < len(s); i++ {
		if rune(s[i]) >= 65 && rune(s[i]) <= 90 || rune(s[i]) >= 97 && rune(s[i]) <= 122 {
			pikkus++
		}
	}
	return pikkus
}
