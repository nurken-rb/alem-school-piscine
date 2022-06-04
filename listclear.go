package piscine

func ListClear(l *List) {
	for i := l.Head; i != nil; i = i.Next {
		l.Head = nil
		l.Tail = nil
	}
}
