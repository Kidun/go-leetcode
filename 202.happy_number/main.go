package isHappy

func isHappy(n int) bool {
	sqMap := map[int]bool{}
	for {
		new_n := 0
		for ; n > 0; n /= 10 {
			new_n += (n % 10) * (n % 10)
		}

		if _, ok := sqMap[new_n]; ok {
			return new_n == 1
		} else {
			sqMap[new_n] = true
		}
		n = new_n
	}
}
