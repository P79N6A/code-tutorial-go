package main
/*
Design a stack that supports push, pop, top, and retrieving the minimum element in constant time.

push(x) -- Push element x onto stack.
pop() -- Removes the element on top of the stack.
top() -- Get the top element.
getMin() -- Retrieve the minimum element in the stack.
Example:
MinStack minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
minStack.getMin();   --> Returns -3.
minStack.pop();
minStack.top();      --> Returns 0.
minStack.getMin();   --> Returns -2.
*/
type MinStack struct {
        lstack []int
        lmin []int
        cap int
        idx int
}


/** initialize your data structure here. */
func Constructor() MinStack {
        cap := 100
        return MinStack{
                lstack:make([]int,100),
                lmin:make([]int,100),
                cap:cap,
                idx:0,
        }
}


func (this *MinStack) Push(x int)  {
        if this.idx >= this.cap-1 {
                nlstack := make([]int,this.cap*2)
                nlmin := make([]int,this.cap*2)
                copy(nlstack,this.lstack)
                copy(nlmin,this.lmin)
                this.lstack=nlstack
                this.lmin=nlmin
                this.cap = this.cap*2
        }
        this.lstack[this.idx] = x
        if this.idx > 0 && this.lmin[this.idx-1] < x {
                this.lmin[this.idx] = this.lmin[this.idx-1]
        } else {
                this.lmin[this.idx]=x
        }
        this.idx += 1
}


func (this *MinStack) Pop()  {
        if this.idx > 0 {
                this.idx -= 1
        }
}


func (this *MinStack) Top() int {
        if this.idx == 0 {return -1}
        return this.lstack[this.idx-1]
}


func (this *MinStack) GetMin() int {
        if this.idx == 0 {return -1}
        return this.lmin[this.idx-1]
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

func main() {
}
