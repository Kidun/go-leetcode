package isPowerOfTwo

func isPowerOfTwo(n int) bool {
	// because 1000 (8) AND 0111 (7) = 0000, and all power of two numbers just 1-bit shifted left
	return (n > 0) && (n&(n-1)) == 0
}
