package main

import "fmt"

/*
Given a linked list, rotate the list to the right by k places, where k is non-negative.

Example 1:

Input: 1->2->3->4->5->NULL, k = 2
Output: 4->5->1->2->3->NULL
Explanation:
rotate 1 steps to the right: 5->1->2->3->4->NULL
rotate 2 steps to the right: 4->5->1->2->3->NULL
Example 2:

Input: 0->1->2->NULL, k = 4
Output: 2->0->1->NULL
Explanation:
rotate 1 steps to the right: 2->0->1->NULL
rotate 2 steps to the right: 1->2->0->NULL
rotate 3 steps to the right: 0->1->2->NULL
rotate 4 steps to the right: 2->0->1->NULL
*/
func printlist(head *ListNode) {
    res := ""
    for head != nil {
        res += fmt.Sprintf("%d", head.Val)
        res += "->"
        head = head.Next
    }
    fmt.Println(res)
}
func main() {
    list := &ListNode{
        Val: 1,
        Next: &ListNode{
            Val: 2,
            Next: &ListNode{
                Val: 3,
                Next: &ListNode{
                    Val: 4,
                    Next: &ListNode{
                        Val: 5,
                    },
                },
            },
        },
    }
    printlist(list)
    h := rotateRight(list, 2)
    printlist(h)

}

/**
 * Definition for singly-linked list.
*/
type ListNode struct {
    Val  int
    Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
    if head == nil {return nil}
    lens := listLen(head)
    k = k % lens
    if k == 0 {return head}
    head, _ = reverse(head, nil)  // reverse all
    p, q := head, head
    for i := 0; i < k; i++ {
        q = p
        p = p.Next
    }
    q.Next,_ = reverse(p, nil)  // reverse last
    p = q.Next   // fix p， tail first pointer
    h, t := reverse(head, p)  // reverse first
    head = h
    t.Next = p // fix tail->next
    return head
}
func reverse(head, tail *ListNode) (*ListNode, *ListNode) {
    if head == nil {
        return nil, nil
    }
    if head.Next == tail {
        return head, head
    }
    h, t := reverse(head.Next, tail)
    t.Next = head
    head.Next = nil
    return h, head
}
func listLen(h *ListNode) int {
    if h == nil {
        return 0
    }
    return 1 + listLen(h.Next)
}
