package missingNumber

func missingNumber(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	// difference between full set n*(n+1)/2 and actual sum
	return ((len(nums) + 1) * len(nums) / 2) - sum
}
