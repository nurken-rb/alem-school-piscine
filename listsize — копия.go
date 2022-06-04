package piscine

func ListSize(l *List) int {
	count := 0
	for i := l.Head; i != nil; i = i.Next {
		count++
	}
	return count
}
