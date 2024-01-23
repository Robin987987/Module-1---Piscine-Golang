package piscine

func SortIntegerTable(table []int) {
	i := 1
	for i < len(table) {
		if table[i-1] > table[i] {
			number := table[i]
			table[i] = table[i-1]
			table[i-1] = number
			i = 1
		} else {
			i++
		}
	}
}
