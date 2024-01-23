package piscine

func ListClear(l *List) {
	iterator := l.Head
	for ; iterator != nil; iterator = iterator.Next {
		l.Head = nil
	}
}
