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

func Test_getIntersectionNode(t *testing.T) {
	list1 := &ListNode{Val: 4}
	list1.Add(1)
	list1.Add(8)
	list1.Add(4)
	list1.Add(5)

	list2 := &ListNode{Val: 5}
	list2.Add(6)
	list2.Add(1)
	list2.Next.Next.Next = list1.Next.Next

	list3 := getIntersectionNode(list1, list2)

	if list3 != list1.Next.Next {
		t.Error("unexpected value for test1")
	}

	list1 = &ListNode{Val: 1}
	list1.Add(9)
	list1.Add(1)
	list1.Add(2)
	list1.Add(4)

	list2 = &ListNode{Val: 3}
	list2.Next = list1.Next.Next.Next

	list3 = getIntersectionNode(list1, list2)

	if list3 != list1.Next.Next.Next {
		t.Error("unexpected value for test1")
	}

	list1 = &ListNode{Val: 2}
	list1.Add(6)
	list1.Add(4)

	list2 = &ListNode{Val: 3}
	list2.Add(5)

	list3 = getIntersectionNode(list1, list2)

	if list3 != nil {
		t.Error("unexpected value for test1")
	}

}
