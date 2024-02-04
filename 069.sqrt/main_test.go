package main

import (
	"strconv"
	"testing"
)

func Test_mySqrt(t *testing.T) {
	x := 195
	res := mySqrt(x)
	if res != 13 {
		t.Error("unexpected value for {" + strconv.Itoa(x) + "}, {" + strconv.Itoa(res) + "}")
	}

	x = 4
	res = mySqrt(x)
	if res != 2 {
		t.Error("unexpected value for {" + strconv.Itoa(x) + "}, {" + strconv.Itoa(res) + "}")
	}

	x = 8
	res = mySqrt(x)
	if res != 2 {
		t.Error("unexpected value for {" + strconv.Itoa(x) + "}, {" + strconv.Itoa(res) + "}")
	}

}
