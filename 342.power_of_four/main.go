package isPowerOfFour

func isPowerOfFour(n int) bool {
	// because 1000 (8) AND 0111 (7) = 0000, and all power of four numbers are just 1-bit shifted left twice
	// 5 = 0101, to check if the "1" is in the even position
	return (n > 0) && (n&(n-1) == 0) && (n&0x55555555 != 0)
}
