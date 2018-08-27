package main

import (
    "container/heap"
    "fmt"
)
/*
Implement FreqStack, a class which simulates the operation of a stack-like data structure.

FreqStack has two functions:

push(int x), which pushes an integer x onto the stack.
pop(), which removes and returns the most frequent element in the stack.
If there is a tie for most frequent element, the element closest to the top of the stack is removed and returned.


Example 1:

Input:
["FreqStack","push","push","push","push","push","push","pop","pop","pop","pop"],
[[],[5],[7],[5],[7],[4],[5],[],[],[],[]]
Output: [null,null,null,null,null,null,null,5,7,5,4]
Explanation:
After making six .push operations, the stack is [5,7,5,7,4,5] from bottom to top.  Then:

pop() -> returns 5, as 5 is the most frequent.
The stack becomes [5,7,5,7,4].

pop() -> returns 7, as 5 and 7 is the most frequent, but 7 is closest to the top.
The stack becomes [5,7,5,4].

pop() -> returns 5.
The stack becomes [5,7,4].

pop() -> returns 4.
The stack becomes [5,7].


Note:

Calls to FreqStack.push(int x) will be such that 0 <= x <= 10^9.
It is guaranteed that FreqStack.pop() won't be called if the stack has zero elements.
The total number of FreqStack.push calls will not exceed 10000 in a single test case.
The total number of FreqStack.pop calls will not exceed 10000 in a single test case.
The total number of FreqStack.push and FreqStack.pop calls will not exceed 150000 across all test cases.
*/

func main() {
    obj := Constructor2()
    obj.Push(5)
    obj.Push(5)
    obj.Push(5)
    fmt.Println(obj.Pop())
}
/*
是个两级stack,先看频次，再看idx
解法一：heap，结构Item包含频次freq和索引idx，Less判断分两步，频次相同看索引，否则看频次。
进入heap元素，重复数字因为freq变化会重复进入。push是O(logn)的，调整堆的复杂度 pop是O(1)的
FreqStack; 再次熟悉下套路的IntHeap写法
解法二: stack of stack。 纵向stack维持的是freq，横向stack维持的idx。这种方式好像更清晰些。
Push/Pop都是O(1)的。重复的数字，还是会重复的进入stack
FreqStack2
*/
type FreqStack struct {
        h    *ItemHeap   // 最大堆
        idx  int    //  因为题目要求在坐标维度，后进先出 stack
        freq map[int]int  // 在频次维度，后进先出 stack
}

func Constructor() FreqStack {
        hh := &ItemHeap{}
        heap.Init(hh)
        return FreqStack{
                h:hh,
                idx:0,
                freq:make(map[int]int),
        }
}

func (this *FreqStack) Push(x int) {
        this.freq[x] += 1
        this.idx += 1
        heap.Push(this.h, item{x, this.freq[x], this.idx})
}

func (this *FreqStack) Pop() int {
        xx := heap.Pop(this.h).(item).num
        //this.idx -= 1 //没有必要缩小,保证谁先谁后就可以了
        this.freq[xx] -= 1
        return xx
}

type item struct {
        /*
        stack num重复的也可以进入;这样就有stack的属性
        */
        num  int
        freq int
        idx  int
}
type ItemHeap []item

func (h ItemHeap)Len() int {
        return len(h)
}
func (h ItemHeap)Swap(i, j int) {
        h[i], h[j] = h[j], h[i]
}
func (h ItemHeap)Less(i, j int) bool {//大顶堆
        hi, hj := h[i], h[j]
        if hi.freq == hj.freq {
                return hi.idx > hj.idx
        }
        return hi.freq > hj.freq
}
func (h *ItemHeap)Push(x interface{}) {
        *h = append(*h, x.(item))
}
func (h *ItemHeap)Pop() interface{} {
        x := (*h)[len(*h) - 1]
        (*h) = (*h)[:len(*h) - 1]
        return x
}
/////////////////////////////////////////////////

type FreqStack2 struct {
    // freq stack. key is num, value is freq
    freq map[int]int
    maxfreq int
    // key is freq, value is a stack.
    group map[int][]int //  index stack
}


func Constructor2() FreqStack2 {
    return FreqStack2{
        freq:make(map[int]int),
        group:make(map[int][]int),
        maxfreq:0,
    }
}


func (this *FreqStack2) Push(x int)  {
    this.freq[x]+=1 // freq add 1
    if this.maxfreq < this.freq[x] {
        this.maxfreq = this.freq[x]
    }
    if len(this.group[this.freq[x]]) <= 0 {
        this.group[this.freq[x]] = make([]int,0)
    }
    // push into stack.
    this.group[this.freq[x]] = append(this.group[this.freq[x]],x)
}


func (this *FreqStack2) Pop() int {
    // 题目说不必考虑pop无效的情况
    xseq := this.group[this.maxfreq]
    x := xseq[len(xseq)-1]
    // pop
    this.group[this.maxfreq] = this.group[this.maxfreq][:len(xseq)-1]
    /*
    想想为什么这个地方不是循环？难道不担心maxfreq-1之后，遇到的是一个空的stack吗？
    答案是不会的，因为maxfreq是一点一点涨上去的，是连续的，比如5,5,5,那最终的stack是
    freq3 : 5
    freq2 : 5
    freq1 : 5
    */
    if len(this.group[this.maxfreq]) <= 0 {
        this.maxfreq -= 1
    }
    this.freq[x]-=1
    return x
}
