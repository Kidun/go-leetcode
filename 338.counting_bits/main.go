package countBits

func countBits(n int) []int {
	// O(N)
	// 1010101 >> 0101010 (which is counted on previous steps + 1 in case of an odd number )
	// 0101010 >> 0010101
	// 0010101 >> 0001010
	// 0001010 >> 0000101
	// 0000101 >> 0000010
	// 0000010 >> 0000001
	bits := make([]int, n+1)
	for i := 0; i <= n; i++ {
		bits[i] = bits[i>>1] + (i & 1)
	}
	return bits
}
