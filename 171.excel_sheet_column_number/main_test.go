package main

import (
	"strconv"
	"testing"
)

func Test_titleToNumber(t *testing.T) {
	a := "A"
	res := titleToNumber(a)
	if res != 1 {
		t.Error("unexpected value for {" + a + "}, {" + strconv.Itoa(res) + "}")
	}

	a = "AB"
	res = titleToNumber(a)
	if res != 28 {
		t.Error("unexpected value for {" + a + "}, {" + strconv.Itoa(res) + "}")
	}

	a = "ZY"
	res = titleToNumber(a)
	if res != 701 {
		t.Error("unexpected value for {" + a + "}, {" + strconv.Itoa(res) + "}")
	}

	a = "ZZY"
	res = titleToNumber(a)
	if res != 18277 {
		t.Error("unexpected value for {" + a + "}, {" + strconv.Itoa(res) + "}")
	}

	a = "CYWOQVJZ"
	res = titleToNumber(a)
	if res != 1234567890*26+26 {
		t.Error("unexpected value for {" + a + "}, {" + strconv.Itoa(res) + "}")
	}
}
