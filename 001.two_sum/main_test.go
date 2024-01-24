package main

import "testing"

func Test_twoSum(t *testing.T) {
	test_res := twoSum([]int{2, 7, 11, 15}, 9)
	if test_res[0] != 0 || test_res[1] != 1 {
		t.Error("unexpected value for {2,7,11,15}, 9", test_res)
	}
	test_res = twoSum([]int{3, 2, 4}, 6)
	if test_res[0] != 1 || test_res[1] != 2 {
		t.Error("unexpected value for {3,2,4}, 6", test_res)
	}
	test_res = twoSum([]int{1, 2, 3, 4, 100, 6, 7, 8, 9, 0}, 100)
	if test_res[0] != 4 || test_res[1] != 9 {
		t.Error("unexpected value for {1,2,3,4,100,6,7,8,9,0}, 100", test_res)
	}
	test_res = twoSum([]int{1, 2, 3, 4, 100, 6, 7, 8, 9, 0}, 200)
	if test_res != nil {
		t.Error("unexpected value for {1, 2, 3, 4, 100, 6, 7, 8, 9, 0}, 200", test_res)
	}
	test_res = twoSum([]int{1}, 1)
	if test_res != nil {
		t.Error("unexpected value for {1}, 1", test_res)
	}
}
