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

func Test_searchInsert(t *testing.T) {
	a := []int{1, 3, 5, 6}
	res := searchInsert(a, 5))
	if res != "2" {
		t.Error("unexpected value for {"+printArray(a, len(a))+"}, |" + strconv.Itoa(res) + "|")
	}

	a = []int{1, 3, 5, 6}
	res = searchInsert(a, 2))
	if res != "1" {
		t.Error("unexpected value for {"+printArray(a, len(a))+"}, |" + strconv.Itoa(res) + "|")
	}
}
