package piscine

func SplitWhiteSpaces(s string) []string {
	TextToString := ""
	t := []string{}
	for i, v := range s {
		if i == lent2(s)-1 && string(v) != " " && string(v) != "\t" && string(v) != "\n" {
			TextToString += string(v)
			t = append(t, TextToString)
		} else if string(v) != " " && string(v) != "\t" && string(v) != "\n" {
			TextToString += string(v)
		} else {
			if len(TextToString) >= 1 {
				t = append(t, TextToString)
			}
			TextToString = ""
		}
	}
	return t
}

func lent2(d string) int {
	inc := 0
	for range d {
		inc++
	}
	return inc
}
