package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	if pos == 0 {
		return l
	}
	count := 1
	i := l
	for ; i != nil; count++ {
		if count == pos {
			return i.Next
		}
		i = i.Next
	}
	return nil
}
