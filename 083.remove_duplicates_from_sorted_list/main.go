package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	/* in-place modification allows us to do this more efficient,
	   but currupts source data
	res := head
	for head.Next != nil {
	    if head.Val == head.Next.Val {
	        head.Next = head.Next.Next
	    } else {
	        head = head.Next
	    }
	}
	return res
	*/
	res := &ListNode{Val: head.Val}
	current := res
	for head.Next != nil {
		if current.Val != head.Next.Val {
			current.Next = &ListNode{Val: head.Next.Val}
			current = current.Next
		}
		head = head.Next
	}
	return res
}

func main() {

}
