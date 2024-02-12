package main

func singleNumber(nums []int) int {
	numsMap := map[int]int{}
	// count all similar values O(n)
	for _, val := range nums {
		numsMap[val]++
	}
	// find out single value O(n)
	for val, count := range numsMap {
		if count == 1 {
			return val
		}
	}
	return 0
}

func main() {

}
