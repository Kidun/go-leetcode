package main

func longestCommonPrefix(strs []string) string {
	if len(strs) > 0 {
		if len(strs) == 1 {
			return (strs[0])
		}
		// it cannot be more than the shortest string
		maxsubstr := strs[0]
		minl := len(strs[0])
		for i := 1; i < len(strs); i++ {
			if len(strs[i]) < minl {
				minl = len(strs[i])
				maxsubstr = strs[i]
			}
		}

		// now check which part of the shortest string matches with other strings
		is_contain := true
		for i := minl; i > 0; i-- {
			substr := maxsubstr[0:i]
			for j := 0; j < len(strs); j++ {
				if substr != strs[j][0:i] {
					// if anything doesn't begin from substr we can skip all other checks
					// and shorten the prefix
					is_contain = false
					break
				}
			}
			if is_contain {
				// since we're here current substr is a prefix for all strings
				return (substr)
			} else {
				// let's shorten the prefix and check again
				is_contain = true
				continue
			}
		}
	}
	return ("")
}

func main() {
}
