package util

// returns the number of digits of a given integer
func LenInt(i int) int {
	if i == 0 {
		return 1
	}

	count := 0
	for i != 0 {
		i /= 10
		count++
	}

	return count
}
