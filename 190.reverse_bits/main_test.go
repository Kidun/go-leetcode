package reverseBits

import (
	"strconv"
	"testing"
)

func Test_reverseBits(t *testing.T) {
	a := uint32(0b00000010100101000001111010011100)
	res := reverseBits(a)
	if res != 964176192 {
		t.Error("unexpected value for {" + strconv.FormatUint(uint64(a), 10) + "}, {" + strconv.FormatUint(uint64(res), 10) + "}")
	}

	a = uint32(0b11111111111111111111111111111101)
	res = reverseBits(a)
	if res != 3221225471 {
		t.Error("unexpected value for {" + strconv.FormatUint(uint64(a), 10) + "}, {" + strconv.FormatUint(uint64(res), 10) + "}")
	}
}
