package piscine

func Join(strs []string, sep string) string {
	s := ""
	for i := range strs {
		s = Concat(s, strs[i])
		if i != len(strs)-1 {
			s = Concat(s, sep)
		}
	}
	return s
}
