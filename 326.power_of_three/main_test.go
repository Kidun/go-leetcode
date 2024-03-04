package isPowerOfThree

import (
	"strconv"
	"testing"
)

func Test_isPowerOfThree(t *testing.T) {
	x := 1
	if !isPowerOfThree(x) {
		t.Error("unexpected value for {" + strconv.Itoa(x) + "}")
	}

	x = 6
	if isPowerOfThree(x) {
		t.Error("unexpected value for {" + strconv.Itoa(x) + "}")
	}

	x = 9
	if !isPowerOfThree(x) {
		t.Error("unexpected value for {" + strconv.Itoa(x) + "}")
	}

	x = -3
	if isPowerOfThree(x) {
		t.Error("unexpected value for {" + strconv.Itoa(x) + "}")
	}

	x = 0
	if isPowerOfThree(x) {
		t.Error("unexpected value for {" + strconv.Itoa(x) + "}")
	}
}
