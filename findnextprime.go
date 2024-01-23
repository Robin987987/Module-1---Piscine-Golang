package piscine

func FindNextPrime(nb int) int {
	if nb < 2 {
		nb = 2
	}
	for ; true; nb++ {
		if IsPrime(nb) {
			return nb
		}
	}
	return nb
}
