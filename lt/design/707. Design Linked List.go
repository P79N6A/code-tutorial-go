package main

import (
    "fmt"
    "strconv"
)

/*
Design your implementation of the linked list. You can choose to use the singly linked list or the doubly linked list. A node in a singly linked list should have two attributes: val and next. val is the value of the current node, and next is a pointer/reference to the next node. If you want to use the doubly linked list, you will need one more attribute prev to indicate the previous node in the linked list. Assume all nodes in the linked list are 0-indexed.

Implement these functions in your linked list class:

get(index) : Get the value of the index-th node in the linked list. If the index is invalid, return -1.
addAtHead(val) : Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list.
addAtTail(val) : Append a node of value val to the last element of the linked list.
addAtIndex(index, val) : Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted.
deleteAtIndex(index) : Delete the index-th node in the linked list, if the index is valid.
Example:

MyLinkedList linkedList = new MyLinkedList();
linkedList.addAtHead(1);
linkedList.addAtTail(3);
linkedList.addAtIndex(1, 2);  // linked list becomes 1->2->3
linkedList.get(1);            // returns 2
linkedList.deleteAtIndex(1);  // now the linked list is 1->3
linkedList.get(1);            // returns 3
Note:

All values will be in the range of [1, 1000].
The number of operations will be in the range of [1, 1000].
Please do not use the built-in LinkedList library.
*/

func main() {
    link := Constructor()
    link.Get(0)
    link.AddAtIndex(1, 2)
    link.Get(0)
    link.Get(1)
    link.AddAtIndex(0, 1)
    link.Get(0)
    link.Get(1)
    print(link.head)
    print(link.tail)
    fmt.Println(link.Get(1))
    fmt.Println(link.n)
    link.DeleteAtIndex(1)
    print(link.head)
    fmt.Println(link.n)
    fmt.Println(link.Get(1))
    fmt.Println(link.Get(2))
    fmt.Println(link.Get(3))
}

type listnode struct {
    val  int
    next *listnode
}
type MyLinkedList struct {
    head *listnode
    tail *listnode
    n    int
}


func get(head *listnode, idx int) int {
    if head == nil {
        return -1
    }
    if idx == 0 {
        return head.val
    }
    return get(head.next, idx-1)
}
func add(head *listnode, idx int, n *listnode) {
    if idx == 0 {
        n.next = head.next
        head.next = n
        return
    }
    add(head.next, idx-1, n)
}
func deletelist(head *listnode, idx int) {
    if head == nil||head.next == nil {return}
    // not first.
    if idx == 0 {
        head.next = head.next.next
        return
    }
    deletelist(head, idx-1)
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
    return MyLinkedList{nil, nil, 0}
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
    if index >= this.n || index<0 {
        return -1
    }
    return get(this.head, index)
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {
    n := &listnode{val, nil}
    if this.head == nil {
        this.head = n
        this.tail = n
    } else {
        n.next = this.head
        this.head = n
    }
    this.n += 1
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
    n := &listnode{val, nil}
    if this.tail == nil {
        this.head = n
        this.tail = n
    } else {
        this.tail.next = n
        this.tail = this.tail.next
    }
    this.n += 1
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {
    if index > this.n || index < 0 {
        return
    }
    if index == this.n {
        this.AddAtTail(val)
        return
    } else if index == 0 {
        this.AddAtHead(val)
        return
    } else {
        n := &listnode{val, nil}
        add(this.head, index-1, n)
        this.n += 1
    }
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
    if index >= this.n || index < 0 {
        return
    }
    if index == 0 {
        this.head = this.head.next
        if this.head == nil {
            this.tail = nil
        }
    } else {
        deletelist(this.head, index-1)
    }
    this.n -= 1
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
