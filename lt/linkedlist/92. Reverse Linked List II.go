package main

import "fmt"

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
    h:=reverseBetween(list,2,4)
    printlist(h)
}
type ListNode struct {
    Val  int
    Next *ListNode
}
func reverseBetween(head *ListNode, m int, n int) *ListNode {
    /*分成三段： head->f2; m1->(m2-1),m2->nil
    判断是否m为0，处理结会不同
    */
    f2 := head
    m1,m2 := head,head
    if m > 1 {
        // m
        for i:=1;i<m-1;i++ {f2 = f2.Next}
        // m
        m1,m2 = f2.Next,f2.Next
    }
    // m+n-m+1=>n+1
    for i:=0;i<=n-m;i++ {m2 = m2.Next}
    mm1,mm2 := reverse(m1,m2)
    mm2.Next = m2
    if m == 1 {return mm1}
    f2.Next = mm1
    return head
}
func reverse(head *ListNode,tail *ListNode) (*ListNode,*ListNode) {
    if head == nil {return nil,nil}
    if head.Next==tail {return head,head} // single node
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