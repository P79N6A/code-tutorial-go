package main

import (
    "strconv"
    "fmt"
)

/*
Sort a linked list in O(n log n) time using constant space complexity.

Example 1:

Input: 4->2->1->3
Output: 1->2->3->4
Example 2:

Input: -1->5->3->4->0
Output: -1->0->3->4->5
*/

func main() {
    h := &ListNode{4,&ListNode{2,&ListNode{1,&ListNode{3,nil}}}}
    //h := &ListNode{4,&ListNode{2,nil}}
    x := sortList(h)
    printlist(x)

}
func printlist(h *ListNode) {
    str := ""
    for h != nil {
        str += strconv.Itoa(h.Val) + "->"
        h = h.Next
    }
    fmt.Println(str)
}

type ListNode struct {
    Val  int
    Next *ListNode
}

func sortList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {return head}
    one,two := head,head
    // 比较容易出问题，比如有两个节点的情况。
    for two != nil {
        two = two.Next
        if two == nil {
            break
        }
        two = two.Next
        if two != nil { // 避免2个节点，one移动两次
            one = one.Next
        }
    }
    n := one.Next
    one.Next = nil
    h1 := sortList(head)
    h2 := sortList(n)
    return mergeOnce(h1,h2)
}
func mergeOnce(h1,h2 *ListNode) *ListNode {
    if h1 == nil {return h2}
    if h2 == nil {return h1}
    if h1.Val < h2.Val {
        h1.Next = mergeOnce(h1.Next,h2)
        return h1
    }
    h2.Next = mergeOnce(h1,h2.Next)
    return h2
}