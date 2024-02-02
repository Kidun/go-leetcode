package main

func lengthOfLastWord(s string) int {
	end_pos := -1
	// iterating the string from the end until we find the end of word
	// then proceed until we find the beginning
	for i := len(s) - 1; i >= 0; i-- {
		if end_pos < 0 {
			if s[i] != ' ' {
				end_pos = i
			}
		} else {
			if s[i] == ' ' {
				return end_pos - i
			}
		}
	}
	return end_pos + 1
}

func main() {

}
