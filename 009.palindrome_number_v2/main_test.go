package main

import "testing"

func Test_isPalindrome(t *testing.T) {

	if isPalindrome(9223372036854775807) {
		t.Error("unexpected value for 9223372036854775807 true")
	}
	if !isPalindrome(12345654321) {
		t.Error("unexpected value for 12345654321 false")
	}
	if isPalindrome(-12345654321) {
		t.Error("unexpected value for -12345654321 true")
	}
	if !isPalindrome(0) {
		t.Error("unexpected value for 0 false")
	}
	if isPalindrome(12345664321) {
		t.Error("unexpected value for 12345664321 true")
	}
	if !isPalindrome(1234554321) {
		t.Error("unexpected value for 1234554321 false")
	}
}
