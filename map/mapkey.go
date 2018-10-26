package main

import "fmt"

type Element struct {
    val int
    name string
}
type Elements []Element

func main() {
    e1 := Element{1,"xx"}
    e2 := Element{1,"xx"}
    e3 := &Element{2,"xx"}
    if e1 == e2 {
        fmt.Println("equall...")
    }
    fmt.Println(e3)
    m := make(map[*Element]bool)
    m[e3]=true
    fmt.Println(m[&e2])
    fmt.Println(m[e3])

}
 type ListNode struct {
         Val int
         Next *ListNode
     }
func removeElements(head *ListNode, val int) *ListNode {
    if head == nil {
        return nil
    }
    head.Next= removeElements(head.Next, val)
    return map[bool]*ListNode{true: head.Next, false: head}[head.Val == val]
}