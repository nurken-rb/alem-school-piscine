package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	asc := true
	desc := true

	for i := 0; i < len(a)-1; i++ {
		if f(a[i], a[i+1]) > 0 {
			asc = false
		}
	}

	for i := 0; i < len(a)-1; i++ {
		if f(a[i], a[i+1]) < 0 {
			desc = false
		}
	}
	return asc || desc
}
