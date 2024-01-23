package piscine

func Index(s string, toFind string) int {
	if len(s) == 0 {
		return -1
	} else if len(toFind) == 0 {
		return 0
	}
	i := 0
	j := 0
	k := 0
	index := 0
	oige := false
	for ; i < len(s); i++ {
		if s[i] == toFind[0] {
			index = i
			j = i
			for ; k < len(toFind); k++ {
				if s[j] == toFind[k] {
					oige = true
				} else {
					oige = false
					break
				}
				j++

			}
		}
		if oige {
			break
		}
	}
	if oige {
		return index
	} else {
		return -1
	}
}
