package main

/*

Reverse a singly linked list.

Example:

Input: 1->2->3->4->5->NULL
Output: 5->4->3->2->1->NULL
Follow up:

A linked list can be reversed either iteratively or recursively. Could you implement both?


*/
/**
 * Definition for singly-linked list.
*/
type ListNode struct {
    Val  int
    Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
    h,_ := reverse(head)
    return h
}
func reverse(head *ListNode) (*ListNode,*ListNode) {
    if head == nil {
        return nil,nil
    }
    if head.Next==nil {
        return head,head
    }
    h,t := reverse(head.Next)
    t.Next = head
    head.Next = nil
    return h,head
}
func main() {

}
