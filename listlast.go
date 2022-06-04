package piscine

func ListLast(l *List) interface{} {
	if l.Head == nil {
		return nil
	}
	i := l.Head
	for i.Next != nil {
		i = i.Next
	}
	return i.Data
}
