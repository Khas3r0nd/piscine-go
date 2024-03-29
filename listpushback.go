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
	n := NodeL{data, nil}
	if l.Head == nil {
		l.Head = &n
		l.Tail = &n
	}
	l.Tail.Next = &n
	l.Tail = &n
}
