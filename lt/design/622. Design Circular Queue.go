package main

import "fmt"

/*
Design your implementation of the circular queue. The circular queue is a linear data structure in which the operations are performed based on FIFO (First In First Out) principle and the last position is connected back to the first position to make a circle. It is also called "Ring Buffer".

One of the benefits of the circular queue is that we can make use of the spaces in front of the queue. In a normal queue, once the queue becomes full, we cannot insert the next element even if there is a space in front of the queue. But using the circular queue, we can use the space to store new values.

Your implementation should support following operations:

MyCircularQueue(k): Constructor, set the size of the queue to be k.
Front: Get the front item from the queue. If the queue is empty, return -1.
Rear: Get the last item from the queue. If the queue is empty, return -1.
enQueue(value): Insert an element into the circular queue. Return true if the operation is successful.
deQueue(): Delete an element from the circular queue. Return true if the operation is successful.
isEmpty(): Checks whether the circular queue is empty or not.
isFull(): Checks whether the circular queue is full or not.


Example:

MyCircularQueue circularQueue = new MycircularQueue(3); // set the size to be 3
circularQueue.enQueue(1);  // return true
circularQueue.enQueue(2);  // return true
circularQueue.enQueue(3);  // return true
circularQueue.enQueue(4);  // return false, the queue is full
circularQueue.Rear();  // return 3
circularQueue.isFull();  // return true
circularQueue.deQueue();  // return true
circularQueue.enQueue(4);  // return true
circularQueue.Rear();  // return 4

Note:

All values will be in the range of [0, 1000].
The number of operations will be in the range of [1, 1000].
Please do not use the built-in Queue library.
*/

func main() {
    queue := Constructor(3)
    fmt.Println(queue.EnQueue(1))
    fmt.Println(queue.EnQueue(2))
    fmt.Println(queue.EnQueue(3))
    fmt.Println(queue.EnQueue(4))
    fmt.Println(queue.queue,queue.head,queue.tail)
    fmt.Println(queue.Rear())
    fmt.Println(queue.IsFull())
    fmt.Println(queue.DeQueue())
    fmt.Println(queue.EnQueue(4))
    fmt.Println(queue.Rear())

}
type MyCircularQueue struct {
    queue []int
    head int // insert at head and head++
    tail int // tail-- then get tail field.
    mod int
}


/** Initialize your data structure here. Set the size of the queue to be k. */
func Constructor(k int) MyCircularQueue {
    return MyCircularQueue{
        queue:make([]int,k+2),
        tail:1,
        head:0,
        mod:k+2,
    }
}
func (this *MyCircularQueue) idx(id int) int {
    if id < 0 {id += this.mod}
    return id%this.mod

}
/** Checks whether the circular queue is empty or not. */
func (this *MyCircularQueue) IsEmpty() bool {
    head := this.idx(this.head+1)
    return head == this.tail
}


/** Checks whether the circular queue is full or not. */
func (this *MyCircularQueue) IsFull() bool {
    tail := this.idx(this.tail+1)
    return tail == this.head

}

/** Insert an element into the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) EnQueue(value int) bool {
    if this.IsFull() {return false}
    this.queue[this.tail]=value
    this.tail=this.idx(this.tail+1)
    return true
}


/** Delete an element from the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) DeQueue() bool {
    if this.IsEmpty() {return false}
    this.head = this.idx(this.head+1)
    return true
}


/** Get the front item from the queue. */
func (this *MyCircularQueue) Front() int {
    if this.IsEmpty() {return -1}
    x := this.idx(this.head+1)
    return this.queue[x]
}


/** Get the last item from the queue. */
func (this *MyCircularQueue) Rear() int {
    if this.IsEmpty() {return -1}
    x := this.idx(this.tail-1)
    return this.queue[x]
}





/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */