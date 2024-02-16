package hammingWeight

func hammingWeight(num uint32) int {
	var res int
	for i := 0; i < 32; i++ {
		res += int(num & 1)
		num >>= 1
		if num == 0 {
			break
		}
	}
	return res
}
