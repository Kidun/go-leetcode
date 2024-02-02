package main

import (
	"strconv"
	"testing"
)

func Test_lengthOfLastWord(t *testing.T) {

	s := "Hello World"
	res := lengthOfLastWord(s)
	if res != 5 {
		t.Error("unexpected value for (" + s + "), result |" + strconv.Itoa(res) + "|")
	}

	s = "   fly me   to   the moon  "
	res = lengthOfLastWord(s)
	if res != 4 {
		t.Error("unexpected value for (" + s + "), result |" + strconv.Itoa(res) + "|")
	}

	s = "luffy is still joyboy"
	res = lengthOfLastWord(s)
	if res != 6 {
		t.Error("unexpected value for (" + s + "), result |" + strconv.Itoa(res) + "|")
	}

	s = "luffyisstilljoyboy"
	res = lengthOfLastWord(s)
	if res != 18 {
		t.Error("unexpected value for (" + s + "), result |" + strconv.Itoa(res) + "|")
	}
}
