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
func Test_largestRectangleArea(t *testing.T) {
	x := []int{2, 1, 5, 6, 2, 3}
	str_x := x.Print()
	res := largestRectangleArea(x)
	if res != 10 {
		t.Error("unexpected value for {" + str_x + "}, {" + strconv.Itoa(res) + "}")
	}

	x = []int{2, 4}
	str_x = x.Print()
	res = largestRectangleArea(x)
	if res != 4 {
		t.Error("unexpected value for {" + str_x + "}, {" + strconv.Itoa(res) + "}")
	}

	x = []int{1, 100, 1000, 0, 10000, 1, 1000, 2000, 3000, 4000, 5000, 4000, 3000, 2000, 1000, 1}
	str_x = x.Print()
	res = largestRectangleArea(x)
	if res != 15000 {
		t.Error("unexpected value for {" + str_x + "}, {" + strconv.Itoa(res) + "}")
	}

	x = []int{0, 0}
	str_x = x.Print()
	res = largestRectangleArea(x)
	if res != 0 {
		t.Error("unexpected value for {" + str_x + "}, {" + strconv.Itoa(res) + "}")
	}

}
