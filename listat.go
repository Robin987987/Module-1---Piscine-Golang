package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	counter := 0
	iterator := l
	for ; counter < pos; iterator = iterator.Next {
		if iterator != nil {
			counter++
		} else {
			return nil
		}
	}
	return iterator
}
