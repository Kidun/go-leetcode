package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	ptrMap := map[*ListNode]bool{}
	for head != nil {
		if _, ok := ptrMap[head]; ok {
			return true
		}
		ptrMap[head] = true
		head = head.Next
	}
	return false
}

func main() {

}
