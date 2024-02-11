package main

func isPalindrome(s string) bool {
	var isAlphanumeric = func(ch byte) bool {
		if (ch >= 65 && ch <= 90) || (ch >= 97 && ch <= 122) || (ch >= 48 && ch <= 57) {
			return true
		}
		return false
	}

	var toLowercase = func(ch byte) byte {
		if ch >= 97 && ch <= 122 {
			return ch - 32
		}
		return ch
	}

	r := len(s) - 1
	for l := 0; l < r; {
		if !isAlphanumeric(s[l]) {
			l++
			continue
		}
		if !isAlphanumeric(s[r]) {
			r--
			continue
		}
		if toLowercase(s[l]) == toLowercase(s[r]) {
			l++
			r--
			continue
		} else {
			return false
		}
	}
	return true
}

func main() {

}
