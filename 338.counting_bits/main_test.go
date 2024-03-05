package countBits

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

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left right);
 */

func Test_countBits(t *testing.T) {
	a := 5
	res := countBits(a)
	str_res := printNumArray(res)
	if str_res != "0 1 1 2 1 2" {
		t.Error("unexpected value for {" + strconv.Itoa(a) + "}  {" + str_res + "}")
	}

	a = 0
	res = countBits(a)
	str_res = printNumArray(res)
	if str_res != "0" {
		t.Error("unexpected value for {" + strconv.Itoa(a) + "}  {" + str_res + "}")
	}

	a = 34
	res = countBits(a)
	str_res = printNumArray(res)
	if str_res != "0 1 1 2 1 2 2 3 1 2 2 3 2 3 3 4 1 2 2 3 2 3 3 4 2 3 3 4 3 4 4 5 1 2 2" {
		t.Error("unexpected value for {" + strconv.Itoa(a) + "}  {" + str_res + "}")
	}
}
