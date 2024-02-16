package reverseBits

func reverseBits(num uint32) uint32 {
	var res uint32
	for i := 0; i < 32; i++ {
		//left shift res and take the last bit
		res = res<<1 | num&1
		num >>= 1
	}
	return res
}
