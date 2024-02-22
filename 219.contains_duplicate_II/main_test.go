package containsNearbyDuplicate

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

func Test_containsNearbyDuplicate(t *testing.T) {
	a := []int{1, 2, 3, 1}
	k := 3
	str_a := printArray(a, len(a))
	if !containsNearbyDuplicate(a, k) {
		t.Error("unexpected value for {" + str_a + "}, {" + strconv.Itoa(k) + "}")
	}

	a = []int{1, 0, 1, 1}
	k = 1
	str_a = printArray(a, len(a))
	if !containsNearbyDuplicate(a, k) {
		t.Error("unexpected value for {" + str_a + "}, {" + strconv.Itoa(k) + "}")
	}

	a = []int{1, 2, 3, 1, 2, 3}
	k = 2
	str_a = printArray(a, len(a))
	if containsNearbyDuplicate(a, k) {
		t.Error("unexpected value for {" + str_a + "}, {" + strconv.Itoa(k) + "}")
	}
}
