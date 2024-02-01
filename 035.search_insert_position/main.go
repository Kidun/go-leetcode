package main

func searchInsert(nums []int, target int) int {
	//sanity check
	if len(nums) == 0 {
		return 0
	}
	if nums[0] >= target {
		return 0
	} else if nums[len(nums)-1] < target {
		return len(nums)
	}
	// as it has O(log n) limitation we have to implement binary search
	// The code can be shorter, I just wanted to make the task a little bit more complicated:
	// insert the value in the beginning of sequence if there is any.
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] >= target && nums[mid-1] < target {
			return mid
		}
		if nums[mid] >= target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

func main() {

}
