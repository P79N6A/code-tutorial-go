package main

import (
    "container/heap"
    "fmt"
)

/*
需要维护的最大的hFixedSize个数字。
所以使用小顶堆
*/
type IntHeap []int
func (h IntHeap)Len()int {return len(h)}
func (h IntHeap)Swap(i,j int) {h[i],h[j]=h[j],h[i]}
func (h IntHeap)Less(i,j int) bool {return h[i]<h[j]} //  小顶堆,这样删除堆顶才是O(1)的，也是Pop干的事情
func (h *IntHeap)Push(x interface{}) {
    *h = append(*h,x.(int))
}
func (h *IntHeap)Pop() interface{} {
    // Pop之前会将第0个和最后一个做交换,所以pop就是固定套路，去掉最后一个，实际上是第一个。
    x := (*h)[len(*h)-1]
    (*h) = (*h)[:len(*h)-1]
    return x
}
const (
    hFixedSize = 6
)
func main() {
    data := []int{1,2,3,4,5,9,8,7,6,5}
    h := &IntHeap{}
    heap.Init(h)

    for i:=0;i<len(data);i++ {
        heap.Push(h,data[i])
        if h.Len() > hFixedSize { // fixed size 超过了就pop
            heap.Pop(h)
        }
    }
    r := make([]int,h.Len())
    for i:=len(r)-1;i>=0;i-- {
        // 注意这个地方是heap.Pop(),这样才是取的队头，并且将0和n-1的元素交换
        // 如果是h.Pop() 则是没有经过swap的队尾了。
        r[i]=heap.Pop(h).(int) // get top N
    }
    fmt.Println(r)
}
