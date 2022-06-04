package piscine

func Swap(a *int, b *int) {
	var temp int = *b
	*b = *a
	*a = temp
}
