package piscine

func IterativePower(nb int, power int) int {
	if power < 0 {
		return 0
	}
	summa := 1
	for ; power > 0; power-- {
		summa *= nb
	}
	return summa
}
