package main

func singleNumber(nums []int) int {
	res := 0
	// using math
	for _, val := range nums {
		// double XOR does nothing to the original value
		// single XOR to 0 gives us the value
		res ^= val
	}

	return res
}

func main() {

}
