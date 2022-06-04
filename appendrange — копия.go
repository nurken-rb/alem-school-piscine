package piscine

func AppendRange(min, max int) []int {
	var arr []int
	if min > max {
		return arr
	}
	for i := 0; i < max-min; i++ {
		arr = append(arr, i+min)
	}
	return arr
}
