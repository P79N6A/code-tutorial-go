package main

import "fmt"

/*
Given a singly linked list L: L0→L1→…→Ln-1→Ln,
reorder it to: L0→Ln→L1→Ln-1→L2→Ln-2→…

You may not modify the values in the list's nodes, only nodes itself may be changed.

Example 1:

Given 1->2->3->4, reorder it to 1->4->2->3.
Example 2:

Given 1->2->3->4->5, reorder it to 1->5->2->4->3.

*/

func main() {
    l := &ListNode{
        1,
        &ListNode{
            2,
            &ListNode{
                3,
                &ListNode{
                    4,
                    &ListNode{
                        5,
                        nil,
                    },
                },
            },
        },
    }
    printlist(l)
    reorderList(l)
    printlist(l)

}
func printlist(h *ListNode)  {
    ret := ""
    for h!= nil {
        ret += fmt.Sprintf("%d,",h.Val)
        h = h.Next
    }
    fmt.Println(ret)
}

/**
 * Definition for singly-linked list.
*/
type ListNode struct {
  Val int
  Next *ListNode
}

func reorderList(head *ListNode)  {
    l := listLen(head)
    if l < 3 {return}
    reorder(head,1,l)
}
func listLen(h *ListNode) int {
    if h == nil {return 0}
    return listLen(h.Next)+1
}

func reorder(h *ListNode, idx, n int) (*ListNode,*ListNode) {
    /*
    递归，第一个返回值是head自己，第二个是head子集合的后续节点
    当剩余1个节点或者两个节点的时候处理返回
    */
    if n%2==0 && idx==n/2 {
        // 剩下两个元素，如何判断剩下几个？跟list长度相关
        tl := h.Next.Next // 记录的后继
        h.Next.Next = nil // 处理完了置空
        return h,tl
    }
    if n%2==1 && idx==(n/2+1)  {
        // 剩下一个元素情况
        tl := h.Next
        h.Next=nil
        return h,tl
    }
    nl,tl := reorder(h.Next,idx+1,n)
    // 这是个从后往前处理的，后序遍历
    tt := tl.Next // 记录后继的下一个
    h.Next=tl // h的下一个是后继
    h.Next.Next = nl // 后继的下一个h.Next 的递归recorder第一个结果
    return h,tt // 返回h，和后继
}
