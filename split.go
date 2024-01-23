package piscine

func Split(s, sep string) []string {
	sona := ""
	uus := 0
	vahele := 0
	for i := range s {
		if s[i] == sep[0] && vahele == 0 {
			for j := range sep {
				if s[i+j] == sep[j] {
					uus = 1
					vahele = len(sep)
				} else {
					uus = 0
					vahele = 0
					break
				}
			}
			if uus == 1 {
				sona += " "
				uus = 0
			}
		}
		if uus == 0 && vahele == 0 {
			sona += string(s[i])
		} else if vahele > 0 {
			vahele--
		} else {
			sona += " "
			// uus = 0
		}
	}
	return SplitWhiteSpaces(sona)
}
