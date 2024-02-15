package main

func majorityElement(nums []int) int {
	numsMap := map[int]int{}
	var res, maxCounter int
	for _, num := range nums {
		numsMap[num]++
		if numsMap[num] > maxCounter {
			maxCounter = numsMap[num]
			res = num
		}
	}
	return res
}

func main() {

}
