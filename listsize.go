package piscine

func ListSize(l *List) int {
	listLength := 0
	if l.Head == nil {
		return 0
	} else {
		for l.Head != nil {
			listLength++
			l.Head = l.Head.Next
		}
	}
	return listLength
}
