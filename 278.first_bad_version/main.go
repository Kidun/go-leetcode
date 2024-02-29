package firstBadVersion

func isBadVersion(version int) bool {
	return version > 123
}

func firstBadVersion(n int) int {
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
