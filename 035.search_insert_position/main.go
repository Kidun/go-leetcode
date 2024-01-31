package main

func searchInsert(nums []int, target int) int {
	//sanity check
	if nums[0] >= target {
		return 0
	} else if nums[len(nums)-1] < target {
		return len(nums)
	}
	// as it has O(log n) limitation we have to implement binary search

	return -1
}

func main() {

}
