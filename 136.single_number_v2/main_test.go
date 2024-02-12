package main

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

func Test_singleNumber(t *testing.T) {
	x := []int{2, 2, 1}
	str_x := printArray(x, len(x))
	res := singleNumber(x)
	if res != 1 {
		t.Error("unexpected value for {" + str_x + "}, {" + strconv.Itoa(res) + "}")
	}

	x = []int{4, 1, 2, 1, 2}
	str_x = printArray(x, len(x))
	res = singleNumber(x)
	if res != 4 {
		t.Error("unexpected value for {" + str_x + "}, {" + strconv.Itoa(res) + "}")
	}

	x = []int{1}
	str_x = printArray(x, len(x))
	res = singleNumber(x)
	if res != 1 {
		t.Error("unexpected value for {" + str_x + "}, {" + strconv.Itoa(res) + "}")
	}
}
