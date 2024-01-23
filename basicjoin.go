package piscine

func BasicJoin(elems []string) string {
	s := ""
	for i := range elems {
		s = Concat(s, elems[i])
	}
	return s
}
