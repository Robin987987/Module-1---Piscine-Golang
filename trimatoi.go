package piscine

func TrimAtoi(s string) int {
	miinus := 1
	a := len(s) - 1
	arv := 0
	for i := 0; i <= a; i++ {
		if arv == 0 && (int(s[i])-45) == 0 {
			miinus = -1
		}
		if IsNumeric(string(s[i])) {
			if i == 0 && (int(s[i])-43) == 0 {
			} else {
				arv = arv*10 + (int(s[i]) - 48)
				if 9 < (int(s[i])-48) || (int(s[i])-48) < 0 {
					return 0
				}
			}
		}
	}
	return arv * miinus
}
