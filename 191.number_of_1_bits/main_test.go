package hammingWeight

import (
	"strconv"
	"testing"
)

func Test_hammingWeight(t *testing.T) {
	a := uint32(0b00000010100101000001111010011100)
	res := hammingWeight(a)
	if res != 12 {
		t.Error("unexpected value for {" + strconv.FormatUint(uint64(a), 2) + "}, {" + strconv.Itoa(res) + "}")
	}

	a = uint32(0b11111111111111111111111111111101)
	res = hammingWeight(a)
	if res != 31 {
		t.Error("unexpected value for {" + strconv.FormatUint(uint64(a), 2) + "}, {" + strconv.Itoa(res) + "}")
	}
}
