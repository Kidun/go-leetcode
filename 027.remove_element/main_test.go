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

func Test_removeElement(t *testing.T) {
	a := []int{3, 2, 2, 3}
	res := printArray(a, removeElement(a, 3))
	if res != "2 2" {
		t.Error("unexpected value for {3, 2, 2, 3}, |" + res + "|")
	}

	a = []int{0, 1, 2, 2, 3, 0, 4, 2}
	res = printArray(a, removeElement(a, 2))
	if res != "0 1 3 0 4" {
		t.Error("unexpected value for {0,1,2,2,3,0,4,2}, |" + res + "|")
	}

	a = []int{}
	res = printArray(a, removeElement(a, 2))
	if res != "Empty slice" {
		t.Error("unexpected value for {}, |" + res + "|")
	}

	a = []int{0, 0, 0, 0, 0}
	res = printArray(a, removeElement(a, 0))
	if res != "Empty slice" {
		t.Error("unexpected value for {0, 0, 0, 0, 0}, |" + res + "|")
	}

	a = []int{0, 0, 1, 0, 0, 0}
	res = printArray(a, removeElement(a, 0))
	if res != "1" {
		t.Error("unexpected value for {0, 0, 1, 0, 0, 0}, |" + res + "|")
	}
}
