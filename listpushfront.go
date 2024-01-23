package piscine

func ListPushFront(l *List, data interface{}) {
	newNode := new(NodeL)
	newNode.Data = data
	if l.Head == nil {
		l.Head = newNode
	} else {
		newNode.Next = l.Head
		l.Head = newNode
	}
}
