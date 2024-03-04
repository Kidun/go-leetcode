package SumRange

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
 * param_1 := obj.SumRange(left,right);
 */

func Test_SumRange(t *testing.T) {
	a := Constructor([]int{-2, 0, 3, -5, 2, -1})
	str_a := printNumArray(a.a)
	if a.SumRange(0, 2) != 1 {
		t.Error("unexpected value for {" + str_a + "}, {" + strconv.Itoa(a.SumRange(0, 2)) + "}")
	}

	if a.SumRange(2, 5) != -1 {
		t.Error("unexpected value for {" + str_a + "}, {" + strconv.Itoa(a.SumRange(2, 5)) + "}")
	}

	if a.SumRange(0, 5) != -3 {
		t.Error("unexpected value for {" + str_a + "}, {" + strconv.Itoa(a.SumRange(0, 5)) + "}")
	}
}
