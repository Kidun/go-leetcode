package main

func climbStairs(n int) int {
	start, next := 0, 1

	// (1) [1 2 3 5 8 13 21 ....] = X(n-1)+X(n-2) i.e. 13 + 21 = 34 for n=8, just a math
	for ; n > 0; n-- {
		start, next = next, start+next
	}
	return next
}

func main() {

}
