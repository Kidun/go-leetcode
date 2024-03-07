package reverseVowels

func reverseVowels(s string) string {
	var isVowel = func(ch rune) bool {
		vowels := map[rune]bool{
			'a': true, 'e': true, 'i': true, 'o': true, 'u': true,
			'A': true, 'E': true, 'I': true, 'O': true, 'U': true}
		_, ok := vowels[ch]
		return ok
	}

	// two pointers moving to each other
	l, r := 0, len(s)-1
	runes := []rune(s)
	for l < r {
		if isVowel(runes[l]) && isVowel(runes[r]) {
			// when both pointers are on vowels, swap them
			runes[l], runes[r] = runes[r], runes[l]
			l++
			r--
		} else if !isVowel(runes[r]) {
			r--
		} else if !isVowel(runes[l]) {
			l++
		}
	}
	return string(runes)
}
