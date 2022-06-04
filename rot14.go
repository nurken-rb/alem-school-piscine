package piscine

func Rot14(s string) string {
	r := []rune(s)
	for i := 0; i < len(r); i++ {
		if r[i] >= 'A' && r[i] <= 'Z' {
			if r[i]+14 > 'Z' {
				r[i] += -12
			} else {
				r[i] += 14
			}
		} else if r[i] >= 'a' && r[i] <= 'z' {
			if r[i]+14 > 'z' {
				r[i] += -12
			} else {
				r[i] += 14
			}
		}
	}
	return string(r)
}
