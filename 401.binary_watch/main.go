package readBinaryWatch

import "fmt"

func readBinaryWatch(turnedOn int) []string {
	// internal function to check given number for number of bits and minutes < 59 and hours < 11
	var trans = func(num int, bits int) string {
		minutes, cnt := 0, 0
		for power := 1; power < 64; power <<= 1 {
			minutes += power * (num & 1)
			if (num & 1) == 1 {
				cnt++
			}
			num >>= 1
			if (cnt > bits) || (minutes > 59) {
				return ""
			}
		}
		hours := 0
		for power := 1; power < 16; power <<= 1 {
			hours += power * (num & 1)
			if (num & 1) == 1 {
				cnt++
			}
			num >>= 1
			if (cnt > bits) || (hours > 11) {
				return ""
			}
		}
		if cnt == bits {
			return fmt.Sprintf("%d:%02d", hours, minutes)
		} else {
			return ""
		}
	}

	res := []string{}
	for i := 0; i <= 763; i++ {
		s := trans(i, turnedOn)
		if s != "" {
			res = append(res, s)
		}
	}
	return res
}
