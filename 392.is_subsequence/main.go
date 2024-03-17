package isSubsequence

func isSubsequence(s string, t string) bool {
	si, ti := 0, 0
	for si = 0; ti < len(t) && si < len(s); {
		if s[si] != t[ti] {
			ti++
			continue
		}
		si++
		ti++
	}
	if ti == len(t) && si < len(s) {
		return false
	}
	return true
}
