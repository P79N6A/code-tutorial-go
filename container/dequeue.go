package main

import (
    "container/list"

    "fmt"
)
type Deque struct {
    list *list.List
}
func (d *Deque)Front() *list.Element{
    return d.list.Front()

}
func (d *Deque)Back() *list.Element{
    return d.list.Back()
}
func (d *Deque)PushFront(v interface{}) {
    d.list.PushFront(v)
}
func (d *Deque)PushBack(v interface{}) {
    d.list.PushBack(v)
}
func (d *Deque)PopFront() interface{} {
    front := d.list.Front()
    return d.list.Remove(front)
}
func (d *Deque)PopBack() interface{} {
    back := d.list.Back()
    return d.list.Remove(back)
}
func (d *Deque)Empty() bool {
    return d.list.Len() == 0
}
func (d *Deque)Dump() {
    for e := d.list.Front();e != nil;e=e.Next() {
        fmt.Printf("%v,",e.Value)
    }
    fmt.Println()
}
func NewDeque() *Deque {
    return &Deque{
        list:list.New(),
    }
}

func main() {
    deque := NewDeque()
    deque.PushFront(1)
    deque.PushFront(2)
    deque.PushFront(3)
    deque.PushFront(4)
    deque.PushBack(-1)
    deque.PushBack(-2)
    deque.PushBack(-3)
    deque.Dump()
    deque.PopBack()
    deque.PopFront()
    deque.Dump()
}
