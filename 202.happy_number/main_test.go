package isHappy

import (
	"strconv"
	"testing"
)

func Test_isHappy(t *testing.T) {
	x := 19
	if !isHappy(x) {
		t.Error("unexpected value for {" + strconv.Itoa(x) + "}")
	}

	x = 2
	if isHappy(x) {
		t.Error("unexpected value for {" + strconv.Itoa(x) + "}")
	}

	x = 1
	if !isHappy(x) {
		t.Error("unexpected value for {" + strconv.Itoa(x) + "}")
	}
}
