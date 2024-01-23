package piscine

func AdvancedSortWordArr(a []string, f func(a, b string) int) {
	temp := ""
	for i := 1; i < len(a); i++ {
		if f(a[i-1], a[i]) == 1 {
			temp = a[i-1]
			a[i-1] = a[i]
			a[i] = temp
			i = 0
		}
	}
}
