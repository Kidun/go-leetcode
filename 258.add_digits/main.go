package addDigits

func addDigits(num int) int {
	// loop approach
	// res := 0
	// for num > 0 {
	// 	res += num % 10
	// 	num /= 10
	// 	if num == 0 {
	// 		if res > 9 {
	// 			num = res
	// 			res = 0
	// 		}
	// 	}
	// }
	// return res

	// math approach
	// 81 -> 9  ----
	// 82 -> 1		|
	// 83 -> 2		|
	// 84 -> 3		|
	// 85 -> 4		|
	// 86 -> 5		|
	// 87 -> 6		|
	// 88 -> 7		|
	// 89 -> 8 -----
	// 90 -> 9
	if num == 0 {
		return 0
	}
	if num%9 == 0 {
		return 9
	}
	return num % 9
}
