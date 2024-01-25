package main

func romanToInt(s string) int {
	// cleaner code, using a map to get roman numerals

	romansMap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	if len(s) < 1 || len(s) > 15 {
		return 0
	}

	res := 0

	// since we have a constraint for characters's set, we can ignore multibyte cases and assume it's a single-byte char/byte array
	for i := 0; i < len(s); i++ {
		if i < len(s)-1 && romansMap[s[i]] < romansMap[s[i+1]] {
			// substract if it's before the higher digit
			res -= romansMap[s[i]]
		} else {
			// add it otherwise
			res += romansMap[s[i]]
		}
	}
	return res
}

func main() {

}
