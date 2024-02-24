package summaryRanges

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

func Test_summaryRanges(t *testing.T) {
	a := []int{0, 1, 2, 4, 5, 7}
	str_a := printNumArray(a)
	res := summaryRanges(a)
	str_res := printStrArray(res)
	if str_res != "0->2,4->5,7" {
		t.Error("unexpected value for {" + str_a + "}, {" + str_res + "}")
	}

	a = []int{0, 2, 3, 4, 6, 8, 9}
	str_a = printNumArray(a)
	res = summaryRanges(a)
	str_res = printStrArray(res)
	if str_res != "0,2->4,6,8->9" {
		t.Error("unexpected value for {" + str_a + "}, {" + str_res + "}")
	}

}
