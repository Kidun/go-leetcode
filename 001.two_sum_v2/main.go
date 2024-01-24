package main

//v2 uses HashMap with O(1) insert/search, one pass loop
func twoSum(nums []int, target int) []int {

	if len(nums) < 2 {
		return nil
	}

	res := []int{0, 0}
	numsMap := map[int]int{}

	var complimentaryValueIndex int
	var isComplimentaryValueInMap bool

	for i := 0; i < len(nums); i++ {
		complimentaryValueIndex, isComplimentaryValueInMap = numsMap[target-nums[i]]
		if isComplimentaryValueInMap && nums[i]+nums[complimentaryValueIndex] == target {
			res[0] = complimentaryValueIndex
			res[1] = i
			return res
		}
		numsMap[nums[i]] = i
	}
	return nil
}

func main() {

}
