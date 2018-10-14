package main

import "fmt"
/*
Implement the following operations of a queue using stacks.

push(x) -- Push element x to the back of queue.
pop() -- Removes the element from in front of queue.
peek() -- Get the front element.
empty() -- Return whether the queue is empty.
Example:

MyQueue queue = new MyQueue();

queue.push(1);
queue.push(2);
queue.peek();  // returns 1
queue.pop();   // returns 1
queue.empty(); // returns false
Notes:

You must use only standard operations of a stack -- which means only push to top, peek/pop from top, size, and is empty operations are valid.
Depending on your language, stack may not be supported natively. You may simulate a stack by using a list or deque (double-ended queue), as long as you use only standard operations of a stack.
You may assume that all operations are valid (for example, no pop or peek operations will be called on an empty queue).
*/

func main() {
        queue := Constructor();

        queue.Push(1);
        queue.Push(2);
        x := queue.Peek();  // returns 1
        fmt.Println(x)
        x = queue.Pop();   // returns 1
        fmt.Println(x)
        queue.Empty(); // returns false
}
type MyQueue struct {
        first,second []int
}


/** Initialize your data structure here. */
func Constructor() MyQueue {
        return MyQueue{
                first:make([]int,0),
                second:make([]int,0),
        }
}


/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int)  {
        this.first = append(this.first,x)
}


/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
        if this.Empty() {
                return -1
        }
        if len(this.second) == 0 {
                for i:=len(this.first)-1;i>=0;i-- {
                        this.second = append(this.second,this.first[i])
                }
                this.first = make([]int,0)
        }
        x := this.second[len(this.second)-1]
        this.second = this.second[:(len(this.second)-1)]
        return x
}


/** Get the front element. */
func (this *MyQueue) Peek() int {
        if this.Empty() {
                return -1
        }
        if len(this.second) == 0 {
                for i:=len(this.first)-1;i>=0;i-- {
                        this.second = append(this.second,this.first[i])
                }
                this.first = make([]int,0)
        }
        x := this.second[len(this.second)-1]
        return x
}


/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
        return len(this.first) + len(this.second) == 0
}


/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */