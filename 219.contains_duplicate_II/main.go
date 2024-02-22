package containsNearbyDuplicate

func containsNearbyDuplicate(nums []int, k int) bool {
	// using hash map we can look for similar values and it costs O(1), total O(N)
	numsMap := map[int]int{}
	for idx, num := range nums {
		if num_idx, ok := numsMap[num]; ok {
			// if it exists, we've found a dup
			if idx-num_idx <= k {
				return true
			}
		}
		// if it's too far we can safely replace it with newly found value or just add it
		numsMap[num] = idx
	}
	return false
}
