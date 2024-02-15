package main

func convertToTitle(columnNumber int) string {
	res := ""
	rest := 0
	var s []int

	// 26 - base, for A-Z numerical system
	for columnNumber > 0 {
		columnNumber, rest = (columnNumber-1)/26, (columnNumber-1)%26
		s = append(s, rest)
	}

	// A = ascii code 65
	for i := len(s) - 1; i >= 0; i-- {
		res = res + string(byte(s[i])+65)
	}

	return res
}

func main() {

}
