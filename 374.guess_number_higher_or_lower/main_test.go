package guessNumber

import (
	"strconv"
	"testing"
)

func Test_guessNumber(t *testing.T) {
	x := 123456
	res := guessNumber(x)
	if res != 30865 {
		t.Error("unexpected value for {" + strconv.Itoa(res) + "}")
	}

}
