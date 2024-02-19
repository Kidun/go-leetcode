package isIsomorphic

import "testing"

func Test_isIsomorphic(t *testing.T) {
	a := "egg"
	b := "add"
	if !isIsomorphic(a, b) {
		t.Error("unexpected value for {" + a + "}, {" + b + "}")
	}

	a = "foo"
	b = "bar"
	if isIsomorphic(a, b) {
		t.Error("unexpected value for {" + a + "}, {" + b + "}")
	}

	a = "paper"
	b = "title"
	if !isIsomorphic(a, b) {
		t.Error("unexpected value for {" + a + "}, {" + b + "}")
	}

	a = "badc"
	b = "baba"
	if isIsomorphic(a, b) {
		t.Error("unexpected value for {" + a + "}, {" + b + "}")
	}

}
