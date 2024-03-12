package guessNumber

// mock function
func guess(num int) int {
	if num < 1234 {
		return 1
	}
	if num == 1234 {
		return 0
	}
	return -1
}

func guessNumber(n int) int {
	// binary search algorithm O(logN)
	l, r := 1, n
	for l < r {
		mid := (l + r) / 2
		g := guess(mid)
		if g == 0 {
			return mid
		}
		if g < 0 {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}
