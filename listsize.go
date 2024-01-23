package piscine

func ListSize(l *List) int {
	counter := 0
	iterator := l.Head
	for ; iterator != nil; iterator = iterator.Next {
		counter++
	}
	return counter
}
