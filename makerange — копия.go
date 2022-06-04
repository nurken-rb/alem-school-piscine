package piscine

func MakeRange(min, max int) []int {
	var arr []int
	if min >= max {
		return arr
	}
	arr = make([]int, max-min)
	for i := 0; i < max-min; i++ {
		arr[i] += i + min
	}
	return arr
}
