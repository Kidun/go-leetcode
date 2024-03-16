package findTheDifference

import (
	"testing"
)

func Test_findTheDifference(t *testing.T) {
	a := "leetcode"
	b := "leettcode"
	res := findTheDifference(a, b)
	if string(res) != "t" {
		t.Error("unexpected value for {" + a + "}, {" + b + "} , {" + string(res) + "}")
	}

}
