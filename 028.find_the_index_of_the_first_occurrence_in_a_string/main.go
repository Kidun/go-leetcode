package main

func strStr(haystack string, needle string) int {
	// sanity check
	if len(needle) > len(haystack) {
		return -1
	}

	// position cannot be more than haystack_end-len(needle)
	for i := 0; i <= len(haystack)-len(needle); i++ {
		//compare substring from position "i"
		for j := 0; j < len(needle); j++ {
			if needle[j] != haystack[i+j] {
				break
			} else if j >= len(needle)-1 {
				return i
			}
		}
	}
	return -1
}

func main() {

}
