package main

func titleToNumber(columnTitle string) int {
	// var pow = func(num int, power int) int {
	// 	if power == 0 {
	// 		return 1
	// 	}
	// 	res := num
	// 	for i := power - 1; i > 0; i-- {
	// 		res *= num
	// 	}
	// 	return res
	// }
	// ascii "A" = 65
	res, base := 0, 1
	for i := len(columnTitle) - 1; i >= 0; i-- {
		res += (int(columnTitle[i]) - 64) * base
		base *= 26
	}
	return res
}
