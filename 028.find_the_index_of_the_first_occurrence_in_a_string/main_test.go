package main

import (
	"strconv"
	"testing"
)

func Test_strStr(t *testing.T) {

	if strStr("sadbutsad", "sad") != 0 {
		t.Error("unexpected value for sadbutsad, sad, |" + strconv.Itoa(strStr("sadbutsad", "sad")) + "|")
	}

	if strStr("leetcode", "leeto") != -1 {
		t.Error("unexpected value for leetcode, leeto, |" + strconv.Itoa(strStr("leetcode", "leeto")) + "|")
	}

	if strStr("leetcodexyz", "xyz") != 8 {
		t.Error("unexpected value for leetcodexyz, xyz, |" + strconv.Itoa(strStr("leetcodexyz", "xyz")) + "|")
	}

}
