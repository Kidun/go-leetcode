package isAnagram2

func isAnagram(s string, t string) bool {
	charMap := map[rune]int{}
	// add runes from s
	for _, ch := range s {
		charMap[ch]++
	}
	// use saved runes for t
	for _, ch := range t {
		charMap[ch]--
	}
	// should be nothing in the map after that
	for _, counter := range charMap {
		if counter != 0 {
			return false
		}
	}
	return true
}
