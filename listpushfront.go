package piscine

func ListPushFront(l *List, data interface{}) {
	n := NodeL{data, nil}
	if l.Head == nil {
		l.Head = &n
		l.Tail = &n
	} else {
		n.Next = l.Head
		l.Head = &n
	}
}
