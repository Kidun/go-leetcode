package reverseList

type ListNode struct {
	Val  int
	Next *ListNode
}

type Stack []*ListNode

// Push adds a value to the top of the stack.
func (st *Stack) Push(node *ListNode) {
	*st = append(*st, node)
}

// Pop removes and returns the top value from the stack.
func (st *Stack) Pop() *ListNode {
	if len(*st) == 0 {
		return nil
	}
	top := (*st)[len(*st)-1]
	*st = (*st)[:len(*st)-1]
	return top
}

// Solution based on LIFO stack
func reverseList(head *ListNode) *ListNode {
	s := Stack{}
	new_head := &ListNode{}
	if head == nil {
		return nil
	}

	for head != nil {
		s.Push(head)
		head = head.Next
	}

	new_head = s.Pop()
	current := new_head
	for len(s) > 0 {
		current.Next = s.Pop()
		current = current.Next
	}
	current.Next = nil
	return new_head
}
