package isPerfectSquare

func isPerfectSquare(num int) bool {
	left := 0
	right := 46340

	// as we can't use math libraries, and brute force approach have O(n),
	// we'll search for the root using binary tree algorithm
	var mid int
	for left < right {
		mid = (right + left) / 2

		if mid*mid >= num {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left*left == num
}
