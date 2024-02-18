package removeElements

import (
	"strconv"
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

func Test_removeElements(t *testing.T) {
	list1 := &ListNode{Val: 1}
	list1.Add(2)
	list1.Add(6)
	list1.Add(3)
	list1.Add(4)
	list1.Add(5)
	list1.Add(6)
	str_list1 := list1.Print()
	a := 6
	res := removeElements(list1, a)
	str_res := res.Print()
	if str_res != "1 2 3 4 5" {
		t.Error("unexpected value for {" + str_list1 + "}, {" + strconv.Itoa(a) + "} = {" + str_res + "}")
	}

	list1 = &ListNode{Val: 7}
	list1.Add(7)
	list1.Add(7)
	list1.Add(7)
	str_list1 = list1.Print()
	a = 7
	res = removeElements(list1, a)
	str_res = res.Print()
	if str_res != "List is empty" {
		t.Error("unexpected value for {" + str_list1 + "}, {" + strconv.Itoa(a) + "} = {" + str_res + "}")
	}

	list1 = &ListNode{Val: 1}
	list1.Add(2)
	list1.Add(2)
	list1.Add(1)
	str_list1 = list1.Print()
	a = 2
	res = removeElements(list1, a)
	str_res = res.Print()
	if str_res != "1 1" {
		t.Error("unexpected value for {" + str_list1 + "}, {" + strconv.Itoa(a) + "} = {" + str_res + "}")
	}
}
