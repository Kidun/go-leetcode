package main

import (
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {
	list1 := &ListNode{Val: 11}
	list1.Add(12)
	list1.Add(13)
	list1.Add(14)
	list1.Add(14)

	list2 := &ListNode{Val: 13}
	list2.Add(13)
	list2.Add(13)

	list3 := mergeTwoLists(list1, list2)
	res := list3.Print()
	if res != "11 12 13 13 13 13 14 14" {
		t.Error("unexpected value for 11 12 13 14 14 + 13 13 13, |" + res + "|")
	}

	list3 = mergeTwoLists(nil, nil)
	res = list3.Print()
	if res != "List is empty" {
		t.Error("unexpected value for nil + nil, |" + res + "|")
	}

	list3 = mergeTwoLists(list1, nil)
	res = list3.Print()
	if res != "11 12 13 14 14" {
		t.Error("unexpected value for 11 12 13 14 14 + nil, |" + res + "|")
	}

	list3 = mergeTwoLists(nil, list2)
	res = list3.Print()
	if res != "13 13 13" {
		t.Error("unexpected value for nil + 13 13 13, |" + res + "|")
	}
}
