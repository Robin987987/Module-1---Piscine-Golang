package piscine

func ConvertBase(nbr, baseFrom, baseTo string) string {
	kümnend := AtoiBase(nbr, baseFrom)
	return NbrBase(kümnend, baseTo)
}
