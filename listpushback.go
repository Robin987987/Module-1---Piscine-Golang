package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushBack(l *List, data interface{}) {
	newNode := new(NodeL)
	newNode.Data = data
	if l.Head == nil {
		l.Head = newNode
	} else {
		iterator := l.Head
		for ; iterator.Next != nil; iterator = iterator.Next {
		}
		iterator.Next = newNode
	}
}
