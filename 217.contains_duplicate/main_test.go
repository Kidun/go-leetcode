package containsDuplicate

import (
	"strconv"
	"testing"
)

// helper function to test slice's content
func printArray(nums []int, nums_len int) string {
	if nums_len <= 0 {
		return "Empty slice"
	}
	res := ""
	for i := 0; i < nums_len; {
		res = res + strconv.Itoa(nums[i])
		i++
		if i != nums_len {
			res = res + " "
		}
	}
	return res
}

func Test_containsDuplicate(t *testing.T) {
	a := []int{1, 2, 3, 1}
	str_a := printArray(a, len(a))
	if !containsDuplicate(a) {
		t.Error("unexpected value for {" + str_a + "}")
	}

	a = []int{1, 2, 3, 4}
	str_a = printArray(a, len(a))
	if containsDuplicate(a) {
		t.Error("unexpected value for {" + str_a + "}")
	}

	a = []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
	str_a = printArray(a, len(a))
	if !containsDuplicate(a) {
		t.Error("unexpected value for {" + str_a + "}")
	}
}
