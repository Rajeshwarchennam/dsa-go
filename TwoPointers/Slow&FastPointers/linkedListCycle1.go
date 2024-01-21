package SlowAndFastPointers

// link-  https://leetcode.com/problems/linked-list-cycle/description/

// Given head, the head of a linked list, determine if the linked list has a cycle in it.

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	slow, fast := head, head
	for {
		if slow != nil {
			slow = slow.Next
		}
		if fast != nil && fast.Next != nil {
			fast = fast.Next.Next
		} else {
			return false
		}
		if slow == fast {
			return true
		}
	}

}
