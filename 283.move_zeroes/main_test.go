package moveZeroes

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

func Test_moveZeroes(t *testing.T) {
	a := []int{0, 1, 0, 3, 12}
	str_a := printNumArray(a)
	moveZeroes(a)
	str_res := printNumArray(a)
	if str_res != "1 3 12 0 0" {
		t.Error("unexpected value for {" + str_a + "}, {" + str_res + "}")
	}

	a = []int{0, 0}
	str_a = printNumArray(a)
	moveZeroes(a)
	str_res = printNumArray(a)
	if str_res != "0 0" {
		t.Error("unexpected value for {" + str_a + "}, {" + str_res + "}")
	}

	a = []int{1, 2}
	str_a = printNumArray(a)
	moveZeroes(a)
	str_res = printNumArray(a)
	if str_res != "1 2" {
		t.Error("unexpected value for {" + str_a + "}, {" + str_res + "}")
	}
}
