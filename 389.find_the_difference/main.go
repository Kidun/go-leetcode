package findTheDifference

func findTheDifference(s string, t string) byte {
	letters := [26]byte{}
	for i := 0; i < len(s); i++ {
		// ascii(a) - 97
		letters[s[i]-97]++
	}
	for i := 0; i < len(t); i++ {
		if letters[t[i]-97] > 0 {
			letters[t[i]-97]--
		} else {
			return t[i]
		}
	}
	return 0
}
