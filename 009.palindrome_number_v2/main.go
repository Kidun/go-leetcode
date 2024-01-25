package main

//constructing reverse number using math, less memory, more cpu (lots of divisions and multiplications)
func isPalindrome(x int) bool {
	res := int(0)

	// if the number is negative or ends with 0 it's not a palindrome number
	if x < 0 || (x != 0 && x%10 == 0) {
		return false
	}

	// it's enough to reverse only the second half of number and remove it from the source.
	for x > res {
		res = res*10 + x%10
		x /= 10
	}

	// If number of digits is odd, it's safe to ignore the middle digit
	return res == x || x == res/10
}

func main() {

}
