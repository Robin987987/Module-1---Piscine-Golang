package piscine

func CountIf(f func(string) bool, tab []string) int {
	count := 0
	for _, item := range tab {
		if f(item) {
			count++
		}
	}
	return count
}
