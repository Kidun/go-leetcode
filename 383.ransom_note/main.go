package canConstruct

func canConstruct(ransomNote string, magazine string) bool {
	// hashmap for O(1) search for available letters
	runeMap := map[rune]int{}

	for _, r := range magazine {
		runeMap[r]++
	}

	for _, r := range ransomNote {
		if runeMap[r] <= 0 {
			return false
		} else {
			runeMap[r]--
		}
	}
	return true
}
