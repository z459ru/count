package count

func Count(s string, r rune) int {
	var cnt int
	for _, l := range s {
		if l == r {
			cnt += 2
		}
	}
	return cnt
}
