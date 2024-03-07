package intersection

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

func Test_intersection(t *testing.T) {
	a := []int{4, 9, 5}
	b := []int{9, 4, 9, 8, 4}
	str_a := printNumArray(a)
	str_b := printNumArray(b)
	res := intersection(a, b)
	str_res := printNumArray(res)
	if (str_res != "9 4") && (str_res != "4 9") {
		t.Error("unexpected value for {" + str_a + "}, {" + str_b + "}, {" + str_res + "}")
	}
}
