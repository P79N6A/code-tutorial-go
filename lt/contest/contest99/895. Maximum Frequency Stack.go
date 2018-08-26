package main

import "container/heap"

func main() {
}

type FreqStack struct {

        h    *ItemHeap
        idx  int
        freq map[int]int
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
