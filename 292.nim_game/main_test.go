package canWinNim

import (
	"strconv"
	"testing"
)

func Test_canWinNim(t *testing.T) {
	a := 1
	if !canWinNim(a) {
		t.Error("unexpected value for {" + strconv.Itoa(a) + "}")
	}

	a = 2
	if !canWinNim(a) {
		t.Error("unexpected value for {" + strconv.Itoa(a) + "}")
	}

	a = 4
	if canWinNim(a) {
		t.Error("unexpected value for {" + strconv.Itoa(a) + "}")
	}
}
