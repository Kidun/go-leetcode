package main

import "testing"

func Test_romanToInt(t *testing.T) {
	test_val := "IV"
	test_res := romanToInt(test_val)
	if test_res != 4 {
		t.Error("unexpected value for ", test_val, " = ", test_res)
	}
	test_val = "VIII"
	test_res = romanToInt(test_val)
	if test_res != 8 {
		t.Error("unexpected value for ", test_val, " = ", test_res)
	}
	test_val = "CMXXXVII"
	test_res = romanToInt(test_val)
	if test_res != 937 {
		t.Error("unexpected value for ", test_val, " = ", test_res)
	}

}
