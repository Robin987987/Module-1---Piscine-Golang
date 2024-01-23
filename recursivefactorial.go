package piscine

func RecursiveFactorial(nb int) int {
	if nb == 0 {
		return 1
	} else if nb > 0 && nb <= 1000 {
		summa := nb * RecursiveFactorial(nb-1)
		return summa
	} else {
		return 0
	}
}
