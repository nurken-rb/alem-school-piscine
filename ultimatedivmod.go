package piscine

func UltimateDivMod(a *int, b *int) {
	var temp int = *a
	*a = *a / *b
	*b = temp % *b
}
