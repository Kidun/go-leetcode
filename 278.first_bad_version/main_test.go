package firstBadVersion

import (
	"strconv"
	"testing"
)

func Test_firstBadVersion(t *testing.T) {
	x := 124
	res := firstBadVersion(x)
	if res != 124 {
		t.Error("unexpected value for {" + strconv.Itoa(x) + "}, {" + strconv.Itoa(res) + "}")
	}

	x = 123456
	res = firstBadVersion(x)
	if res != 124 {
		t.Error("unexpected value for {" + strconv.Itoa(x) + "}, {" + strconv.Itoa(res) + "}")
	}
}
