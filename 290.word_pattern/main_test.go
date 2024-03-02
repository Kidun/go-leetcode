package wordPattern

import "testing"

func Test_wordPattern(t *testing.T) {
	a := "abba"
	b := "dog cat cat dog"
	if !wordPattern(a, b) {
		t.Error("unexpected value for {" + a + "}, {" + b + "}")
	}

	a = "abba"
	b = "dog cat cat fish"
	if wordPattern(a, b) {
		t.Error("unexpected value for {" + a + "}, {" + b + "}")
	}

	a = "aaaa"
	b = "dog cat cat dog"
	if wordPattern(a, b) {
		t.Error("unexpected value for {" + a + "}, {" + b + "}")
	}

	a = "abba"
	b = "cat cat cat cat"
	if wordPattern(a, b) {
		t.Error("unexpected value for {" + a + "}, {" + b + "}")
	}

	a = "aaaaba"
	b = "cat cat cat cat"
	if wordPattern(a, b) {
		t.Error("unexpected value for {" + a + "}, {" + b + "}")
	}

	a = "aaa"
	b = "cat cat cat cat"
	if wordPattern(a, b) {
		t.Error("unexpected value for {" + a + "}, {" + b + "}")
	}
}
