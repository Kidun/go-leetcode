package toHex

import (
	"strconv"
	"testing"
)

func Test_toHex(t *testing.T) {
	a := 26
	b := toHex(a)
	if b != "1a" {
		t.Error("unexpected value for {" + strconv.Itoa(a) + "}, {" + b + "}")
	}

	a = -1
	b = toHex(a)
	if b != "ffffffff" {
		t.Error("unexpected value for {" + strconv.Itoa(a) + "}, {" + b + "}")
	}

	a = 0
	b = toHex(a)
	if b != "0" {
		t.Error("unexpected value for {" + strconv.Itoa(a) + "}, {" + b + "}")
	}

}
