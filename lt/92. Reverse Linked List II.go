package main

import "fmt"

/*

Reverse a linked list from position m to n. Do it in one-pass.

Note: 1 ≤ m ≤ n ≤ length of list.

Example:

Input: 1->2->3->4->5->NULL, m = 2, n = 4
Output: 1->4->3->2->5->NULL
*/
/**
 * Definition for singly-linked list.
*/
type ListNode struct {
    Val  int
    Next *ListNode
}
func reverseBetween(head *ListNode, m int, n int) *ListNode {
    if n < m {return head}
    s := head
    var sp *ListNode
    for i:=0;i<m-1&&s!=nil;i++ {
        sp=s
        s = s.Next
    }
    e := head
    for i:=0;i<n&&e!=nil;i++ {
        e = e.Next
    }
    h,t := reverse(s,e)
    if t != nil {
        t.Next=e
    }
    if sp == nil {
        return h
    }
    sp.Next=h
    return head
}
func reverse(head *ListNode,tail *ListNode) (*ListNode,*ListNode) {
    if head == nil {
        return nil,nil
    }
    if head.Next==tail {
        return head,head
    }
    h,t := reverse(head.Next,tail)
    t.Next = head
    head.Next = nil
    return h,head
}
func printlist(head *ListNode) {
    res := ""
    for head != nil {
        res += fmt.Sprintf("%d",head.Val)
        res += "->"
        head = head.Next
    }
    fmt.Println(res)
}
func main() {
    list := &ListNode{
        Val:1,
        Next:&ListNode{
            Val:2,
            Next:&ListNode{
                Val:3,
                Next:&ListNode{
                    Val:4,
                    Next:&ListNode{
                        Val:5,
                    },
                },
            },
        },
    }
    printlist(list)
    h:=reverseBetween(list,2,3)
    printlist(h)
    //l2 := &ListNode{1,nil}
    //h1:=reverseBetween(l2,1,1)
    //printlist(h1)
}
