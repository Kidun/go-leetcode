package isIsomorphic

func isIsomorphic(s string, t string) bool {
	letterMap := map[byte]byte{}
	// it's also possible to double check 1. s->t 2. t->s
	// or make reverse map to check double booking
	letterMapRev := map[byte]byte{}
	for i := 0; i < len(s); i++ {
		if letter, ok := letterMap[s[i]]; ok {
			if letter == t[i] {
				continue
			} else {
				return false
			}
		} else {
			if _, ok := letterMapRev[t[i]]; ok {
				return false
			}
			letterMap[s[i]] = t[i]
			letterMapRev[t[i]] = s[i]
		}
	}
	return true
}
