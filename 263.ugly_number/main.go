package isUgly

func isUgly(n int) bool {
	// edge cases
	if n <= 0 {
		return false
	}
	var divn = func(n int, d int) int {
		for n%d == 0 {
			n /= d
		}
		return n
	}
	n = divn(n, 2)
	n = divn(n, 3)
	n = divn(n, 5)
	// after all divisions of ugly number there must be reminder == 1
	return n == 1
}
