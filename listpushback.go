package piscine

func ListPushBack(l *List, data interface{}) {
	n := &NodeL{Data: data}

	if l.Head == nil {
		l.Head = n
		return
	}

	for i := l.Head; i != nil; i = i.Next {
		if i.Next == nil {
			i.Next = n
			return
		}
	}
}
