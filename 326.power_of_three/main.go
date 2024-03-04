package isPowerOfThree

func isPowerOfThree(n int) bool {
	// 2^31=2147483648
	// 3^20=3486784401
	// so all 3^n that less than 2^31 will divide 2^20 without a reminder
	if n > 0 {
		return n == 1 || 3486784401%n == 0
	}
	return false
}
