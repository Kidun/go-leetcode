package main

func removeDuplicates(nums []int) int {
	// sanity check
	if len(nums) <= 0 {
		return 0
	}

	new_index := 0
	for index := 1; index < len(nums); index++ {
		// add in-place only values that are bigger than previous values
		if nums[index] > nums[new_index] {
			new_index++
			nums[new_index] = nums[index]
		}
	}
	return new_index + 1
}

func main() {

}
