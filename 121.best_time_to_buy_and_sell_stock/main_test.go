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

func Test_maxProfit(t *testing.T) {
	a := []int{7, 2, 4, 8, 1}
	str_a := printArray(a, len(a))
	res := maxProfit(a)
	if res != 6 {
		t.Error("unexpected value for {" + str_a + "} => {" + strconv.Itoa(res) + "}")
	}

	a = []int{7, 1, 5, 3, 6, 4}
	str_a = printArray(a, len(a))
	res = maxProfit(a)
	if res != 5 {
		t.Error("unexpected value for {" + str_a + "} => {" + strconv.Itoa(res) + "}")
	}

	a = []int{7, 6, 4, 3, 1}
	str_a = printArray(a, len(a))
	res = maxProfit(a)
	if res != 0 {
		t.Error("unexpected value for {" + str_a + "} => {" + strconv.Itoa(res) + "}")
	}
}
