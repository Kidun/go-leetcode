package readBinaryWatch

import (
	"strconv"
	"testing"
)

// helper function to test slice's content
func printNumArray(slice []int) string {
	if len(slice) <= 0 {
		return "Empty slice"
	}
	res := ""
	for i := 0; i < len(slice); {
		res = res + strconv.Itoa(slice[i])
		i++
		if i != len(slice) {
			res = res + " "
		}
	}
	return res
}

func printStrArray(slice []string) string {
	if len(slice) <= 0 {
		return "Empty slice"
	}
	res := ""
	for i := 0; i < len(slice); {
		res = res + slice[i]
		i++
		if i != len(slice) {
			res = res + ","
		}
	}
	return res
}

func Test_readBinaryWatch(t *testing.T) {
	a := 1
	res := readBinaryWatch(a)
	str_res := printStrArray(res)
	if str_res != "0:01,0:02,0:04,0:08,0:16,0:32,1:00,2:00,4:00,8:00" {
		t.Error("unexpected value for {" + strconv.Itoa(a) + "}, {" + str_res + "}")
	}

	a = 9
	res = readBinaryWatch(a)
	str_res = printStrArray(res)
	if str_res != "Empty slice" {
		t.Error("unexpected value for {" + strconv.Itoa(a) + "}, {" + str_res + "}")
	}

	a = 0
	res = readBinaryWatch(a)
	str_res = printStrArray(res)
	if str_res != "0:00" {
		t.Error("unexpected value for {" + strconv.Itoa(a) + "}, {" + str_res + "}")
	}

	a = 8
	res = readBinaryWatch(a)
	str_res = printStrArray(res)
	if str_res != "7:31,7:47,7:55,7:59,11:31,11:47,11:55,11:59" {
		t.Error("unexpected value for {" + strconv.Itoa(a) + "}, {" + str_res + "}")
	}

}
