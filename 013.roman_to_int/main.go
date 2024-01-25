package main

func romanToInt(s string) int {
	// direct approach hard to read code, but quick and memory-efficient
	if len(s) < 1 || len(s) > 15 {
		return 0
	}

	res := 0

	// since we have a constraint for characters's set, we can ignore multibyte cases and assume it's a single-byte char array
	for i := 0; i < len(s); i++ {
		switch {
		case s[i] == 'I':
			if i < len(s)-1 && s[i+1] == 'V' {
				res += 4
				i++
			} else if i < len(s)-1 && s[i+1] == 'X' {
				res += 9
				i++
			} else {
				res += 1
			}
		case s[i] == 'V':
			res += 5
		case s[i] == 'X':
			if i < len(s)-1 && s[i+1] == 'L' {
				res += 40
				i++
			} else if i < len(s)-1 && s[i+1] == 'C' {
				res += 90
				i++
			} else {
				res += 10
			}
		case s[i] == 'L':
			res += 50
		case s[i] == 'C':
			if i < len(s)-1 && s[i+1] == 'D' {
				res += 400
				i++
			} else if i < len(s)-1 && s[i+1] == 'M' {
				res += 900
				i++
			} else {
				res += 100
			}
		case s[i] == 'D':
			res += 500
		case s[i] == 'M':
			res += 1000
		}
	}
	return res
}

func main() {

}
