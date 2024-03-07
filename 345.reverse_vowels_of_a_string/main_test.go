package reverseVowels

import "testing"

func Test_reverseVowels(t *testing.T) {
	a := "hello"
	b := reverseVowels(a)
	if b != "holle" {
		t.Error("unexpected value for {" + a + "}, {" + b + "}")
	}

	a = "a."
	b = reverseVowels(a)
	if b != "a." {
		t.Error("unexpected value for {" + a + "}, {" + b + "}")
	}

	a = "leetcode"
	b = reverseVowels(a)
	if b != "leotcede" {
		t.Error("unexpected value for {" + a + "}, {" + b + "}")
	}

	a = "alemmutttiiIcoOqEqUr"
	b = reverseVowels(a)
	if b != "UlEmmOtttoIiciuqeqar" {
		t.Error("unexpected value for {" + a + "}, {" + b + "}")
	}

}
