package piscine

func Any(f func(string) bool, a []string) bool {
	for _, item := range a {
		if f(item) {
			return true
		}
	}
	return false
}
