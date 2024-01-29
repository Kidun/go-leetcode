package main

import (
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// Add a new value to the list
func (list *ListNode) Add(val int) {
	new_node := &ListNode{Val: val}

	current := list
	for current.Next != nil {
		current = current.Next
	}
	current.Next = new_node
}

// Print the list, for debug and testing purpose
func (list *ListNode) Print() string {
	current := list
	res := ""

	if current == nil {
		res = "List is empty"
		return res
	} else {
		for current != nil {
			res = res + strconv.Itoa(current.Val)
			current = current.Next
			if current != nil {
				res = res + " "
			}
		}
		return res
	}
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// sanity checks
	if list1 == nil && list2 == nil {
		return nil
	} else if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	}

	var list3, list3_head *ListNode = nil, nil

	for list1 != nil || list2 != nil {
		// create a list and save a head pointer on the first iteration
		if list3 != nil {
			list3.Next = &ListNode{}
			list3 = list3.Next
		} else {
			list3 = &ListNode{}
			list3_head = list3
		}

		// add the largest number and move pointer to the next node until both lists are processed
		// we need to add two additional checks for nil to avoid invalid memory reference
		if list2 == nil || (list1 != nil && list1.Val <= list2.Val) {
			list3.Val = list1.Val
			list1 = list1.Next
		} else {
			list3.Val = list2.Val
			list2 = list2.Next
		}
	}

	return list3_head
}

func main() {

}
