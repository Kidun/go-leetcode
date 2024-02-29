package firstBadVersion

func isBadVersion(version int) bool {
	return version > 123
}

func firstBadVersion(n int) int {
	left := 1
	right := n
	var cur int
	for left <= right {
		cur = (left + right) / 2
		if isBadVersion(cur) {
			right = cur - 1
		} else {
			left = cur + 1
		}
	}
	return cur
}
