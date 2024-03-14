package canConstruct

import "testing"

func Test_canConstruct(t *testing.T) {
	a := "hello"
	b := "holle"
	if !canConstruct(a, b) {
		t.Error("unexpected value for {" + a + "}, {" + b + "}")
	}

	a = "leetcode"
	b = "leotcedexc"
	if !canConstruct(a, b) {
		t.Error("unexpected value for {" + a + "}, {" + b + "}")
	}

	a = "ee"
	b = "e"
	if canConstruct(a, b) {
		t.Error("unexpected value for {" + a + "}, {" + b + "}")
	}

}
