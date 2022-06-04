package piscine

func BasicJoin(elems []string) string {
	var s string
	for i := 0; i < len(elems); i++ {
		s += elems[i]
	}
	return s
}
