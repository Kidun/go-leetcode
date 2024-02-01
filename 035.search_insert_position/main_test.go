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
	a := []int{1, 1, 1, 1, 1, 3, 3, 3, 4, 4, 4, 4, 4, 4, 4, 5, 6}
	ins := 3
	res := searchInsert(a, ins)
	if res != 5 {
		t.Error("unexpected value for inserting " + strconv.Itoa(ins) + " into {" + printArray(a, len(a)) + "}, result is |" + strconv.Itoa(res) + "|")
	}

	a = []int{1, 1, 1, 1, 3, 3, 3, 4, 4, 4, 4, 4, 4, 4, 5, 6}
	res = searchInsert(a, ins)
	if res != 4 {
		t.Error("unexpected value for inserting " + strconv.Itoa(ins) + " into {" + printArray(a, len(a)) + "}, result is |" + strconv.Itoa(res) + "|")
	}

	a = []int{1, 1, 1, 3, 3, 3, 4, 4, 4, 4, 4, 4, 4, 5, 6}
	res = searchInsert(a, ins)
	if res != 3 {
		t.Error("unexpected value for inserting " + strconv.Itoa(ins) + " into {" + printArray(a, len(a)) + "}, result is |" + strconv.Itoa(res) + "|")
	}

	a = []int{1, 1, 3, 3, 3, 4, 4, 4, 4, 4, 4, 4, 5, 6}
	res = searchInsert(a, ins)
	if res != 2 {
		t.Error("unexpected value for inserting " + strconv.Itoa(ins) + " into {" + printArray(a, len(a)) + "}, result is |" + strconv.Itoa(res) + "|")
	}

	a = []int{1, 3, 3, 3, 4, 4, 4, 4, 4, 4, 4, 5, 6}
	res = searchInsert(a, ins)
	if res != 1 {
		t.Error("unexpected value for inserting " + strconv.Itoa(ins) + " into {" + printArray(a, len(a)) + "}, result is |" + strconv.Itoa(res) + "|")
	}

	a = []int{1, 3, 3, 3, 4, 4, 4, 4, 4, 4, 4, 5, 6}
	ins = 2
	res = searchInsert(a, ins)
	if res != 1 {
		t.Error("unexpected value for inserting " + strconv.Itoa(ins) + " into {" + printArray(a, len(a)) + "}, result is |" + strconv.Itoa(res) + "|")
	}

	a = []int{1, 3, 5, 6}
	ins = 5
	res = searchInsert(a, ins)
	if res != 2 {
		t.Error("unexpected value for inserting " + strconv.Itoa(ins) + " into {" + printArray(a, len(a)) + "}, result is |" + strconv.Itoa(res) + "|")
	}

	ins = 2
	res = searchInsert(a, ins)
	if res != 1 {
		t.Error("unexpected value for {" + printArray(a, len(a)) + "}, |" + strconv.Itoa(res) + "|")
	}

	ins = 7
	res = searchInsert(a, ins)
	if res != 4 {
		t.Error("unexpected value for {" + printArray(a, len(a)) + "}, |" + strconv.Itoa(res) + "|")
	}

	ins = 0
	res = searchInsert(a, ins)
	if res != 0 {
		t.Error("unexpected value for {" + printArray(a, len(a)) + "}, |" + strconv.Itoa(res) + "|")
	}
}
