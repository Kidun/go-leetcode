package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	// using hashmap we'll get O(N) solution with clean code
	ptrMap := map[*ListNode]bool{}
	// save all pointers from listA to a hashmap
	for headA != nil {
		ptrMap[headA] = true
		headA = headA.Next
	}
	// if we find the same pointer in the hashmap as in listB
	// return immidiately, as it's a head of the intersection
	for headB != nil {
		if _, ok := ptrMap[headB]; ok {
			return headB
		}
		headB = headB.Next
	}

	return nil
}

func main() {

}
