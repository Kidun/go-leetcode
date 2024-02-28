package isUgly

import (
	"strconv"
	"testing"
)

func Test_isUgly(t *testing.T) {
	x := 1
	if !isUgly(x) {
		t.Error("unexpected value for {" + strconv.Itoa(x) + "}")
	}

	x = 6
	if !isUgly(x) {
		t.Error("unexpected value for {" + strconv.Itoa(x) + "}")
	}

	x = 14
	if isUgly(x) {
		t.Error("unexpected value for {" + strconv.Itoa(x) + "}")
	}

	x = 0
	if isUgly(x) {
		t.Error("unexpected value for {" + strconv.Itoa(x) + "}")
	}
}
