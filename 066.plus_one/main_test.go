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

func Test_plusOne(t *testing.T) {
	src := []int{1, 2, 3}
	str_src := printArray(src, len(src))
	res := plusOne(src)
	str_res := printArray(res, len(res))
	if res != "1 2 4" {
		t.Error("unexpected value for {" + str_src + "}, {" + str_res + "}")
	}

	src = []int{9, 9, 9, 9}
	str_src = printArray(src, len(src))
	res = plusOne(src)
	str_res = printArray(res, len(res))
	if res != "1 0 0 0 0" {
		t.Error("unexpected value for {" + str_src + "}, {" + str_res + "}")
	}

	src = []int{4, 3, 2, 1}
	str_src = printArray(src, len(src))
	res = plusOne(src)
	str_res = printArray(res, len(res))
	if res != "4 3 2 2" {
		t.Error("unexpected value for {" + str_src + "}, {" + str_res + "}")
	}

	src = []int{9}
	str_src = printArray(src, len(src))
	res = plusOne(src)
	str_res = printArray(res, len(res))
	if res != "1 0" {
		t.Error("unexpected value for {" + str_src + "}, {" + str_res + "}")
	}
}
