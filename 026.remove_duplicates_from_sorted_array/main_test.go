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

func Test_removeDuplicates(t *testing.T) {
	a := []int{1, 2, 2, 2, 3, 3, 5}
	res := printArray(a, removeDuplicates(a))
	if res != "1 2 3 5" {
		t.Error("unexpected value for 1 2 2 2 3 3 5, |" + res + "|")
	}

	a = []int{1, 1, 1, 2, 2, 3, 4, 4, 5, 5}
	res = printArray(a, removeDuplicates(a))
	if res != "1 2 3 4 5" {
		t.Error("unexpected value for 1 1 1 2 2 3 4 4 5 5, |" + res + "|")
	}

	a = []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	res = printArray(a, removeDuplicates(a))
	if res != "1" {
		t.Error("unexpected value for 1 1 1 1 1 1 1 1 1 1, |" + res + "|")
	}

	a = []int{}
	res = printArray(a, removeDuplicates(a))
	if res != "Empty slice" {
		t.Error("unexpected value for empty slice, |" + res + "|")
	}

	a = []int{1, 4}
	res = printArray(a, removeDuplicates(a))
	if res != "1 4" {
		t.Error("unexpected value for 1 4, |" + res + "|")
	}

	a = []int{1}
	res = printArray(a, removeDuplicates(a))
	if res != "1" {
		t.Error("unexpected value for 1, |" + res + "|")
	}
}
