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

func Test_majorityElement(t *testing.T) {
	a := []int{3, 2, 3}
	str_a := printArray(a, len(a))
	res := majorityElement(a)
	if res != 3 {
		t.Error("unexpected value for {" + str_a + "} => {" + strconv.Itoa(res) + "}")
	}

	a = []int{2, 2, 1, 1, 1, 2, 2}
	str_a = printArray(a, len(a))
	res = majorityElement(a)
	if res != 2 {
		t.Error("unexpected value for {" + str_a + "} => {" + strconv.Itoa(res) + "}")
	}
}
