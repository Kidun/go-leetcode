package containsDuplicate

func containsDuplicate(nums []int) bool {
	// using hash map we can look for similar values and it costs O(1), total O(N)
	numsMap := map[int]bool{}
	for _, num := range nums {
		if numsMap[num] {
			// if it exists, we've found a dup
			return true
		}
		// add new value to a hashmap
		numsMap[num] = true
	}
	return false
}
