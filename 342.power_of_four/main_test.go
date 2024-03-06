package isPowerOfFour

import (
	"strconv"
	"testing"
)

func Test_isPowerOfFour(t *testing.T) {
	x := 1
	if !isPowerOfFour(x) {
		t.Error("unexpected value for {" + strconv.Itoa(x) + "}")
	}

	x = 4096
	if !isPowerOfFour(x) {
		t.Error("unexpected value for {" + strconv.Itoa(x) + "}")
	}

	x = -4096
	if isPowerOfFour(x) {
		t.Error("unexpected value for {" + strconv.Itoa(x) + "}")
	}

	x = 6
	if isPowerOfFour(x) {
		t.Error("unexpected value for {" + strconv.Itoa(x) + "}")
	}

	x = 64
	if !isPowerOfFour(x) {
		t.Error("unexpected value for {" + strconv.Itoa(x) + "}")
	}
}
