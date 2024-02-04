package main

import (
	"strconv"
	"testing"
)

func Test_climbStairs(t *testing.T) {
	x := 2
	res := climbStairs(x)
	if res != 2 {
		t.Error("unexpected value for {" + strconv.Itoa(x) + "}, {" + strconv.Itoa(res) + "}")
	}

	x = 3
	res = climbStairs(x)
	if res != 3 {
		t.Error("unexpected value for {" + strconv.Itoa(x) + "}, {" + strconv.Itoa(res) + "}")
	}
}
