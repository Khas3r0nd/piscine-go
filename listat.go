package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	counter := 0
	var answer *NodeL
	if l == nil {
		return nil
	}
	for counter != pos {
		answer = l
		l = l.Next
		counter++
	}
	return answer
}
