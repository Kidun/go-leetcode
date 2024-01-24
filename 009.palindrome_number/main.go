package main

// direct approach comparing digits using array
func isPalindrome(x int) bool {
	var nums [19]int8
	var numsLength int8 = 0

	// negative numbers are not palindrome numbers
	if x < 0 {
		return false
	}
	
	if x != 0 {
		// conversion int to an array of decimal digits
		for x != 0 {
			nums[numsLength] = int8(x % 10)
			x /= 10
			numsLength++
		}
	} else {
		numsLength = 1
	}

	// checking if array is a palindrome
	for i := int8(0); i <= numsLength/2; i++ {
		if nums[i] != nums[numsLength-i-1] {
			return false
		}
	}
	return true
}

func main() {

}
