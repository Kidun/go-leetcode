package main

import "testing"

func Test_isValid(t *testing.T) {
	if isValid("{}[]([()({()})])") != true {
		t.Error("unexpected value for {}[]([()({()})])")
	}
	if isValid("") != true {
		t.Error("unexpected value for \"\"")
	}
	if isValid("]") != false {
		t.Error("unexpected value for ]")
	}
	if isValid("}") != false {
		t.Error("unexpected value for }")
	}
	if isValid(")") != false {
		t.Error("unexpected value for )")
	}
	if isValid("[") != false {
		t.Error("unexpected value for [")
	}
	if isValid("{") != false {
		t.Error("unexpected value for {")
	}
	if isValid("(") != false {
		t.Error("unexpected value for (")
	}
	if isValid("{[()()(])}") != false {
		t.Error("unexpected value for {[()()(])}")
	}

}
