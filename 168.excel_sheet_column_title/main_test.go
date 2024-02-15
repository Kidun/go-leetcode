package main

import (
	"strconv"
	"testing"
)

func Test_convertToTitle(t *testing.T) {
	a := 1
	res := convertToTitle(a)
	if res != "A" {
		t.Error("unexpected value for " + strconv.Itoa(a) + "{" + res + "}")
	}

	a = 28
	res = convertToTitle(a)
	if res != "AB" {
		t.Error("unexpected value for " + strconv.Itoa(a) + "{" + res + "}")
	}

	a = 701
	res = convertToTitle(a)
	if res != "ZY" {
		t.Error("unexpected value for " + strconv.Itoa(a) + "{" + res + "}")
	}

	a = 18277
	res = convertToTitle(a)
	if res != "ZZY" {
		t.Error("unexpected value for " + strconv.Itoa(a) + "{" + res + "}")
	}

	a = 1234567890*26 + 26
	res = convertToTitle(a)
	if res != "CYWOQVJZ" {
		t.Error("unexpected value for " + strconv.Itoa(a) + "{" + res + "}")
	}
}
