package main

import (
	"testing"
)

func Test_plusOne(t *testing.T) {
	stra := "11"
	strb := "1"
	res := addBinary(stra, strb)
	if res != "100" {
		t.Error("unexpected value for {" + stra + " + " + strb + "}, {" + res + "}")
	}

	stra = "1010"
	strb = "1011"
	res = addBinary(stra, strb)
	if res != "10101" {
		t.Error("unexpected value for {" + stra + " + " + strb + "}, {" + res + "}")
	}

	stra = "0"
	strb = "11111111"
	res = addBinary(stra, strb)
	if res != "11111111" {
		t.Error("unexpected value for {" + stra + " + " + strb + "}, {" + res + "}")
	}

	stra = "1111"
	strb = "1111"
	res = addBinary(stra, strb)
	if res != "11110" {
		t.Error("unexpected value for {" + stra + " + " + strb + "}, {" + res + "}")
	}

}
