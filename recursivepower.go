package piscine

func RecursivePower(nb int, power int) int {
	var summa int
	if power < 0 {
		return 0
	} else if power == 0 {
		return 1
	}
	summa = nb * RecursivePower(nb, power-1)
	return summa
}
