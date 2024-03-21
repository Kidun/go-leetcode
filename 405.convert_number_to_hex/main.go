package toHex

func toHex(num int) string {
	//edge case that's not covered by loop
	if num == 0 {
		return "0"
	}
	chars := "0123456789abcdef"
	res := ""
	unum := uint32(num)
	for unum != 0 {
		// 4-bit shift right is equal to division by 2^4=16
		n := unum & 0xf
		unum >>= 4
		res = string(chars[n]) + res
	}
	return res
}
