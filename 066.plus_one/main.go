package main

func plusOne(digits []int) []int {
	over := false
	new_len := len(digits)
	if digits[len(digits)-1] == 9 {
		over = true
		new_len = len(digits) + 1
		for i := 0; i < len(digits); i++ {
			if digits[i] != 9 {
				new_len = len(digits)
				break
			}
		}
	}
	res := make([]int, new_len)
	if over && new_len > len(digits) {
		res[len(digits)] = 0
	}
	for i := len(digits) - 1; i >= 0; i-- {
		if over {
			if digits[i] == 9 {
				if i > 0 {
					res[i] = 0
				} else {
					res[i] = 1
				}
			} else {
				over = false
				res[i] = digits[i] + 1
			}
		} else {
			if i == len(digits)-1 {
				res[i] = digits[i] + 1
			} else {
				res[i] = digits[i]
			}
		}
	}
	return res
}

func main() {

}
