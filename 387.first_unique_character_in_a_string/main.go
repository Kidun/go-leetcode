package firstUniqChar

func firstUniqChar(s string) int {
	// unicode friendly solution, if it consists of only latin letters
	// it's more optimal to use [26]int array and store there counters
	sMap := map[rune]int{}
	sRune := []rune(s)
	for _, r := range sRune {
		sMap[r]++
	}
	for n, r := range sRune {
		if sMap[r] == 1 {
			return n
		}
	}
	return -1
}
