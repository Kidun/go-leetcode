package main

func addBinary(a string, b string) string {
	num := len(b)
	res := ""
	if len(a) >= len(b) {
		num = len(a)
	}
	index := 0
	over := false
	for i := num; i > 0; i-- {
		var aa, bb byte
		if len(a)-index-1 < 0 {
			aa = '0'
		} else {
			aa = a[len(a)-index-1]
		}
		if len(b)-index-1 < 0 {
			bb = '0'
		} else {
			bb = b[len(b)-index-1]
		}

		if aa == '0' && bb == '0' {
			if over {
				res = "1" + res
				over = false
			} else {
				res = "0" + res
			}
		} else if (aa == '1' && bb == '0') || (aa == '0' && bb == '1') {
			if over {
				res = "0" + res
			} else {
				res = "1" + res
			}
		} else if aa == '1' && bb == '1' {
			if over {
				res = "1" + res
			} else {
				res = "0" + res
				over = true
			}
		}
		index++
	}
	if over {
		res = "1" + res
	}
	return res
}

func main() {

}
