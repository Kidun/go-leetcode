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

	x = 4
	res = climbStairs(x)
	if res != 5 {
		t.Error("unexpected value for {" + strconv.Itoa(x) + "}, {" + strconv.Itoa(res) + "}")
	}

	x = 5
	res = climbStairs(x)
	if res != 8 {
		t.Error("unexpected value for {" + strconv.Itoa(x) + "}, {" + strconv.Itoa(res) + "}")
	}

	x = 6
	res = climbStairs(x)
	if res != 13 {
		t.Error("unexpected value for {" + strconv.Itoa(x) + "}, {" + strconv.Itoa(res) + "}")
	}

	x = 1
	res = climbStairs(x)
	if res != 1 {
		t.Error("unexpected value for {" + strconv.Itoa(x) + "}, {" + strconv.Itoa(res) + "}")
	}

}
