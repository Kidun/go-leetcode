package isSubsequence

import "testing"

func Test_isSubsequence(t *testing.T) {
	a := "abc"
	b := "ahbgdc"
	if !isSubsequence(a, b) {
		t.Error("unexpected value for {" + a + "}, {" + b + "}")
	}

	a = "axc"
	b = "ahbgdc"
	if isSubsequence(a, b) {
		t.Error("unexpected value for {" + a + "}, {" + b + "}")
	}

	a = "abc"
	b = "abb"
	if isSubsequence(a, b) {
		t.Error("unexpected value for {" + a + "}, {" + b + "}")
	}

}
