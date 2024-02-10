package main

import "testing"

func Test_isPalindrome(t *testing.T) {
	a := "A man, a plan, a canal: Panama"
	if isPalindrome(a) != true {
		t.Error("unexpected value for {" + a + "}")
	}

	a = "race a car"
	if isPalindrome(a) == true {
		t.Error("unexpected value for {" + a + "}")
	}

	a = " "
	if isPalindrome(a) != true {
		t.Error("unexpected value for {" + a + "}")
	}
}
