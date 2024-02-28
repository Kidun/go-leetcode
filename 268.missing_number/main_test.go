package missingNumber

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

func Test_summaryRanges(t *testing.T) {
	a := []int{9, 6, 4, 2, 3, 5, 7, 0, 1}
	str_a := printNumArray(a)
	res := missingNumber(a)
	if res != 8 {
		t.Error("unexpected value for {" + str_a + "}, {" + strconv.Itoa(res) + "}")
	}

	a = []int{3, 0, 1}
	str_a = printNumArray(a)
	res = missingNumber(a)
	if res != 2 {
		t.Error("unexpected value for {" + str_a + "}, {" + strconv.Itoa(res) + "}")
	}

	a = []int{0, 1}
	str_a = printNumArray(a)
	res = missingNumber(a)
	if res != 2 {
		t.Error("unexpected value for {" + str_a + "}, {" + strconv.Itoa(res) + "}")
	}

}
