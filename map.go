package piscine

func Map(f func(int) bool, a []int) []bool {
	result := make([]bool, len(a))

	for i := 0; i < len(a); i++ {
		isF := f(a[i])
		result[i] = isF
	}
	return result
}
