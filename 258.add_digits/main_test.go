package addDigits

import (
	"strconv"
	"testing"
)

func Test_addDigits(t *testing.T) {
	a := 38
	res := addDigits(a)
	if res != 2 {
		t.Error("unexpected value for {" + strconv.Itoa(a) + "}, {" + strconv.Itoa(res) + "}")
	}

	a = 0
	res = addDigits(a)
	if res != 0 {
		t.Error("unexpected value for {" + strconv.Itoa(a) + "}, {" + strconv.Itoa(res) + "}")
	}
}
