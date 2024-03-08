package isPerfectSquare

import (
	"strconv"
	"testing"
)

func Test_isPerfectSquare(t *testing.T) {
	x := 195
	if isPerfectSquare(x) {
		t.Error("unexpected value for {" + strconv.Itoa(x) + "}")
	}

	x = 225
	if !isPerfectSquare(x) {
		t.Error("unexpected value for {" + strconv.Itoa(x) + "}")
	}

	x = 4
	if !isPerfectSquare(x) {
		t.Error("unexpected value for {" + strconv.Itoa(x) + "}")
	}

	x = 2147395600
	if !isPerfectSquare(x) {
		t.Error("unexpected value for {" + strconv.Itoa(x) + "}")
	}
}
