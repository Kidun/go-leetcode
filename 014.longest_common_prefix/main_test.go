package main

import "testing"

func Test_longestCommonPrefix(t *testing.T) {
	test_val := []string{"abc", "abcd", "ab", "abcdeeeee"}
	test_res := longestCommonPrefix(test_val)
	if test_res != "ab" {
		t.Error("unexpected value for ", test_val, " = ", test_res)
	}

	test_val2 := []string{"abcd", "abc", "ab", "a", "qqq"}
	test_res2 := longestCommonPrefix(test_val2)
	if test_res2 != "" {
		t.Error("unexpected value for ", test_val2, " = ", test_res2)
	}

	test_val3 := []string{"abcd"}
	test_res3 := longestCommonPrefix(test_val3)
	if test_res3 != "abcd" {
		t.Error("unexpected value for ", test_val3, " = ", test_res3)
	}

	test_val4 := []string{}
	test_res4 := longestCommonPrefix(test_val4)
	if test_res4 != "" {
		t.Error("unexpected value for ", test_val4, " = ", test_res4)
	}

	test_val5 := []string{"abcdeeeee", "abcdeeee1", "abcdeeeee", "abcdeeeee"}
	test_res5 := longestCommonPrefix(test_val5)
	if test_res5 != "abcdeeee" {
		t.Error("unexpected value for ", test_val5, " = ", test_res5)
	}

}
