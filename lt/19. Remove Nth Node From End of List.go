package main
/*
Given linked list: 1->2->3->4->5, and n = 2.

After removing the second node from the end, the linked list becomes 1->2->3->5.

Note:
Given n will always be valid.
Try to do this in one pass.
*/
type ListNode struct {
        Val int
        Next *ListNode
}
func removeNthFromEnd(head *ListNode, n int) *ListNode {
        p,q := head,head
        for i:=0;i<n;i++ {
                q = q.Next
        }
        if q == nil {
                return head.Next
        }
        for {
                q = q.Next
                if q == nil {
                        p.Next = p.Next.Next
                        break
                }
                p = p.Next
        }
        return head
}
func main() {
}
