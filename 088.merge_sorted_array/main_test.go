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

func Test_merge(t *testing.T) {
	a := []int{1, 2, 3, 0, 0, 0}
	m := 3
	b := []int{2, 5, 6}
	n := 3
	str_a := printArray(a, m)
	str_b := printArray(b, n)
	merge(a, m, b, n)
	str_res := printArray(a, m+n)
	if str_res != "1 2 2 3 5 6" {
		t.Error("unexpected value for {" + str_a + "}, {" + str_b + "} => {" + str_res + "}")
	}

	a = []int{4, 5, 6, 0, 0, 0, 0, 0, 0}
	m = 3
	b = []int{1, 2, 6, 7, 8, 9}
	n = 6
	str_a = printArray(a, m)
	str_b = printArray(b, n)
	merge(a, m, b, n)
	str_res = printArray(a, m+n)
	if str_res != "1 2 4 5 6 6 7 8 9" {
		t.Error("unexpected value for {" + str_a + "}, {" + str_b + "} => {" + str_res + "}")
	}

	a = []int{4, 5, 6, 0, 0, 0}
	m = 3
	b = []int{1, 2, 3}
	n = 3
	str_a = printArray(a, m)
	str_b = printArray(b, n)
	merge(a, m, b, n)
	str_res = printArray(a, m+n)
	if str_res != "1 2 3 4 5 6" {
		t.Error("unexpected value for {" + str_a + "}, {" + str_b + "} => {" + str_res + "}")
	}

	a = []int{1}
	m = 1
	b = []int{}
	n = 0
	str_a = printArray(a, m)
	str_b = printArray(b, n)
	merge(a, m, b, n)
	str_res = printArray(a, m+n)
	if str_res != "1" {
		t.Error("unexpected value for {" + str_a + "}, {" + str_b + "} => {" + str_res + "}")
	}

	a = []int{0}
	m = 0
	b = []int{1}
	n = 1
	str_a = printArray(a, m)
	str_b = printArray(b, n)
	merge(a, m, b, n)
	str_res = printArray(a, m+n)
	if str_res != "1" {
		t.Error("unexpected value for {" + str_a + "}, {" + str_b + "} => {" + str_res + "}")
	}
}
