package firstBadVersion

// mock function
func isBadVersion(version int) bool {
	return version > 123
}

func firstBadVersion(n int) int {
	// O(log(n)) binary search
	left := 1
	for left < n {
		mid := (left + n) / 2
		if isBadVersion(mid) {
			n = mid
		} else {
			left = mid + 1
		}
	}
	return left
}
