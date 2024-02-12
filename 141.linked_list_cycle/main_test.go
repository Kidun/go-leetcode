package main

import (
	"testing"
)

// Add a new value to the list
func (list *ListNode) Add(val int) {
	new_node := &ListNode{Val: val}

	current := list
	for current.Next != nil {
		current = current.Next
	}
	current.Next = new_node
}

func Test_hasCycle(t *testing.T) {
	list1 := &ListNode{Val: 3}
	list1.Add(2)
	list1.Add(0)
	list1.Add(-4)
	list1.Next.Next.Next = list1.Next

	if !hasCycle(list1) {
		t.Error("unexpected value for test1")
	}

	list1 = &ListNode{Val: 1}
	list1.Add(2)
	list1.Next.Next = list1.Next
	if !hasCycle(list1) {
		t.Error("unexpected value for test2")
	}

	list1 = &ListNode{Val: 1}
	if hasCycle(list1) {
		t.Error("unexpected value for test3")
	}

	if hasCycle(nil) {
		t.Error("unexpected value for test4")
	}

	list1 = &ListNode{Val: 1}
	list1.Add(2)
	if hasCycle(list1) {
		t.Error("unexpected value for test5")
	}

	list1 = &ListNode{Val: 1}
	list1.Next = list1
	if !hasCycle(list1) {
		t.Error("unexpected value for test3")
	}
}
