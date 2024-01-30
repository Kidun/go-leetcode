package main

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	new_index := 0
	for index := 0; index < len(nums); index++ {
		if nums[index] != val {
			nums[new_index] = nums[index]
			new_index++
		}
	}
	return new_index
}

func main() {

}
