package isAnagram

import "testing"

func Test_isAnagram(t *testing.T) {
	a := "anagram"
	b := "nagaram"
	if !isAnagram(a, b) {
		t.Error("unexpected value for {" + a + "}, {" + b + "}")
	}

	a = "rat"
	b = "car"
	if isAnagram(a, b) {
		t.Error("unexpected value for {" + a + "}, {" + b + "}")
	}

	a = "aa"
	b = "bb"
	if isAnagram(a, b) {
		t.Error("unexpected value for {" + a + "}, {" + b + "}")
	}

	a = "hqbqo"
	b = "lsnma"
	if isAnagram(a, b) {
		t.Error("unexpected value for {" + a + "}, {" + b + "}")
	}
}
