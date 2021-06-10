package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// dummy node 哑节点很重要
type ListNode struct {
	Val  int
	Next *ListNode
}

// 一趟删除
func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{0, head}
	first := head
	second := dummy
	for i := 0; i < n; i++ {
		first = first.Next
	}

	for first.Next != nil {
		first = first.Next
		second = second.Next
	}
	second.Next = second.Next.Next
	return dummy.Next

	//dummy := &ListNode{0, head}
	//first, second := head, dummy
	//for i := 0; i < n; i++ {
	//	first = first.Next
	//}
	//for ; first != nil; first = first.Next {
	//	second = second.Next
	//}
	//second.Next = second.Next.Next
	//return dummy.Next
}
