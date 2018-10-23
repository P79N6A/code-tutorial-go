package main

import "fmt"

/*
Design your implementation of the circular double-ended queue (deque).

Your implementation should support following operations:

MyCircularDeque(k): Constructor, set the size of the deque to be k.
insertFront(): Adds an item at the front of Deque. Return true if the operation is successful.
insertLast(): Adds an item at the rear of Deque. Return true if the operation is successful.
deleteFront(): Deletes an item from the front of Deque. Return true if the operation is successful.
deleteLast(): Deletes an item from the rear of Deque. Return true if the operation is successful.
getFront(): Gets the front item from the Deque. If the deque is empty, return -1.
getRear(): Gets the last item from Deque. If the deque is empty, return -1.
isEmpty(): Checks whether Deque is empty or not.
isFull(): Checks whether Deque is full or not.


Example:

MyCircularDeque circularDeque = new MycircularDeque(3); // set the size to be 3
circularDeque.insertLast(1);			// return true
circularDeque.insertLast(2);			// return true
circularDeque.insertFront(3);			// return true
circularDeque.insertFront(4);			// return false, the queue is full
circularDeque.getRear();  			// return 2
circularDeque.isFull();				// return true
circularDeque.deleteLast();			// return true
circularDeque.insertFront(4);			// return true
circularDeque.getFront();			// return 4


Note:

All values will be in the range of [0, 1000].
The number of operations will be in the range of [1, 1000].
Please do not use the built-in Deque library.
*/

func main() {
    circularDeque := Constructor(3); // set the size to be 3
    fmt.Println(circularDeque.InsertLast(1));			// return true
    fmt.Println(circularDeque.InsertLast(2));			// return true
    fmt.Println(circularDeque.InsertFront(3));			// return true
    fmt.Println(circularDeque.InsertFront(4));			// return false, the queue is full
    fmt.Println(circularDeque.GetRear());  			// return 2
    fmt.Println(circularDeque.IsFull());				// return true
    fmt.Println(circularDeque.DeleteLast());			// return true
    fmt.Println(circularDeque.InsertFront(4));			// return true
    fmt.Println(circularDeque.GetFront());			// return 4
}
type MyCircularDeque struct {
    q []int
    start,end int
    k int
}


/** Initialize your data structure here. Set the size of the deque to be k. */
func Constructor(k int) MyCircularDeque {
    return MyCircularDeque{
        //浪费2个空间做标记，start,end不存数据
        q:make([]int,k+2),
        start:0,
        end:1,
        k:k+2,
    }

}

/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertFront(value int) bool {
    if this.IsFull() {return false}
    this.q[this.start]=value
    this.start-=1
    if this.start < 0 {this.start += this.k}
    this.start %=this.k
    return true
}


/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertLast(value int) bool {
    if this.IsFull() {return false}
    this.q[this.end]=value
    this.end +=1
    this.end %=this.k
    return true
}


/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteFront() bool {
    if this.IsEmpty() {return false}
    this.start+=1
    this.start %=this.k
    return true
}


/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteLast() bool {
    if this.IsEmpty() {return false}
    this.end -=1
    if this.end<0 {this.end+=this.k}
    this.end %=this.k
    return true
}


/** Get the front item from the deque. */
func (this *MyCircularDeque) GetFront() int {
    if this.IsEmpty() {return -1}
    x := this.start
    x += 1
    x %=this.k
    return this.q[x]
}


/** Get the last item from the deque. */
func (this *MyCircularDeque) GetRear() int {
    if this.IsEmpty() {return -1}
    x := this.end
    x -= 1
    if x < 0 {x += this.k}
    x %=this.k
    return this.q[x]
}


/** Checks whether the circular deque is empty or not. */
func (this *MyCircularDeque) IsEmpty() bool {
    x := this.start
    x += 1
    x %=this.k
    return x == this.end
}


/** Checks whether the circular deque is full or not. */
func (this *MyCircularDeque) IsFull() bool {
    x := this.start
    x -= 1
    if x < 0 {x += this.k}
    x %=this.k
    return x == this.end
}


/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */