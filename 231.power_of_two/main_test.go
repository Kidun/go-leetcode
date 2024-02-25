package isPowerOfTwo

import (
	"strconv"
	"testing"
)

func Test_isPowerOfTwo(t *testing.T) {
	x := 1
	if !isPowerOfTwo(x) {
		t.Error("unexpected value for {" + strconv.Itoa(x) + "}")
	}

	x = 4096
	if !isPowerOfTwo(x) {
		t.Error("unexpected value for {" + strconv.Itoa(x) + "}")
	}

	x = -4096
	if isPowerOfTwo(x) {
		t.Error("unexpected value for {" + strconv.Itoa(x) + "}")
	}

	x = 7
	if isPowerOfTwo(x) {
		t.Error("unexpected value for {" + strconv.Itoa(x) + "}")
	}

}
