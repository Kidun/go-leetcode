package firstUniqChar

import (
	"strconv"
	"testing"
)

func Test_firstUniqChar(t *testing.T) {
	a := "leetcode"
	res := firstUniqChar(a)
	if res != 0 {
		t.Error("unexpected value for {" + a + "}, {" + strconv.Itoa(res) + "}")
	}

	a = "loveleetcode"
	res = firstUniqChar(a)
	if res != 2 {
		t.Error("unexpected value for {" + a + "}, {" + strconv.Itoa(res) + "}")
	}

	a = "aabb"
	res = firstUniqChar(a)
	if res != -1 {
		t.Error("unexpected value for {" + a + "}, {" + strconv.Itoa(res) + "}")
	}
}
