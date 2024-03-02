package wordPattern

func wordPattern(pattern string, s string) bool {
	wordnum := 0
	word := ""
	strMap := map[byte]string{}
	patMap := map[string]byte{}
	for i := 0; i <= len(s); i++ {
		if i < len(s) && string(s[i]) != " " {
			word += string(s[i])
		} else {
			// case when patterb has less letters than there are words
			if wordnum >= len(pattern) {
				return false
			}
			w, wok := strMap[pattern[wordnum]]
			pat, pok := patMap[word]
			// if word doesn't follow the pattern
			if (wok && w != word) || (pok && pat != pattern[wordnum]) {
				return false
			} else {
				strMap[pattern[wordnum]] = word
				patMap[word] = pattern[wordnum]
			}
			word = ""
			wordnum++
		}
	}
	// condition to cover case when pattern has more letters than there are words
	return wordnum == len(pattern)
}
