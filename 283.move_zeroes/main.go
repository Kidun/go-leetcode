package moveZeroes

func moveZeroes(nums []int) {
	pos := 0
	for _, num := range nums {
		if num != 0 {
			nums[pos] = num
			pos++
		}
	}
	for ; pos < len(nums); pos++ {
		nums[pos] = 0
	}
}
