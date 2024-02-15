package main

func majorityElement(nums []int) int {
	// hashmap solution
	// numsMap := map[int]int{}
	// half := len(nums) / 2
	res, maxCounter := nums[0], 1
	// numsMap[nums[0]] = 1
	for i := 1; i < len(nums); i++ {
		// numsMap[nums[i]]++
		// if numsMap[nums[i]] > maxCounter {
		// 	maxCounter, res = numsMap[nums[i]], nums[i]
		// 	if maxCounter > half {
		// 		return res
		// 	}
		// }

		// solution that exploits >⌊n/2⌋ condition
		// Boyer-Moore Majority Vote Algorithm
		if maxCounter == 0 {
			res, maxCounter = nums[i], 1
		} else {
			if nums[i] == res {
				maxCounter++
			} else {
				maxCounter--
			}
		}
	}

	return res
}

func main() {

}
