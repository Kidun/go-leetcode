package main

func singleNumber(nums []int) int {
	numsMap := map[int]int{}
	for _, val := range nums {
		_, ok := numsMap[val]
		if ok {
			delete(numsMap, val)
		} else {
			numsMap[val] = 1
		}
	}
	for val, count := range numsMap {
		if count == 1 {
			return val
		}
	}
	return 0
}

func main() {

}
