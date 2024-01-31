package main

func strStr(haystack string, needle string) int {
	// sanity check
	if len(needle) > len(haystack) {
		return -1
	}

	// position cannot be more than haystack_end-len(needle)
	for i := 0; i <= len(haystack)-len(needle); i++ {
		if haystack[i:i+len(needle)] == needle {
			return i
		}
	}
	return -1
}

func main() {

}
