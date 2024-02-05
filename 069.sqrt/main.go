package main

func mySqrt(x int) int {
	// sanity checks, border conditions
	left := 0
	right := 46340
	if x >= 2147395600 {
		return right
	}

	// as we can't use math libraries, and brute force approach have O(n),
	// we'll search for the root using binary tree algorithm
	var mid int
	for left <= right {
		mid = (right + left) / 2

		if mid*mid > x {
			right = mid
		} else {
			left = mid
			if (mid+1)*(mid+1) > x {
				//return the value if we've found that x lies between mid^2 and (mid+1)^2
				return mid
			}
		}
	}
	return mid
}

func main() {

}
