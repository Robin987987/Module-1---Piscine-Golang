package piscine

func SortWordArr(a []string) {
	temp := ""
	for i := 1; i < len(a); i++ {
		if a[i-1] > a[i] {
			temp = a[i-1]
			a[i-1] = a[i]
			a[i] = temp
			i = 0
		}
	}
}
