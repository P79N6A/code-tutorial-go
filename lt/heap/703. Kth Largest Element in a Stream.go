package main

import (
    "container/heap"
    "fmt"
)

/*
Design a class to find the kth largest element in a stream. Note that it is the kth largest element in the sorted order, not the kth distinct element.

Your KthLargest class will have a constructor which accepts an integer k and an integer array nums, which contains initial elements from the stream. For each call to the method KthLargest.add, return the element representing the kth largest element in the stream.

Example:

int k = 3;
int[] arr = [4,5,8,2];
KthLargest kthLargest = new KthLargest(3, arr);
kthLargest.add(3);   // returns 4
kthLargest.add(5);   // returns 5
kthLargest.add(10);  // returns 5
kthLargest.add(9);   // returns 8
kthLargest.add(4);   // returns 8
Note:
You may assume that nums' length ≥ k-1 and k ≥ 1.


*/

func main() {
    obj := Constructor(3,[]int{4,5,8,2})
    fmt.Println(obj.Add(3))
    fmt.Println(obj.Add(5))
    fmt.Println(obj.Add(10))
    fmt.Println(obj.Add(9))
    fmt.Println(obj.Add(4))

}
////////////////////////

////////////////////////
type IntHeap []int
func (h *IntHeap)Len() int {return len(*h)}
func (h *IntHeap)Swap(i,j int) {(*h)[i],(*h)[j]=(*h)[j],(*h)[i]}
func (h *IntHeap)Less(i,j int) bool {return (*h)[i]<(*h)[j]}
func (h *IntHeap)Push(x interface{}) {
    *h = append(*h,x.(int))
}
func (h *IntHeap)Top() int {
    return (*h)[0]
}
func (h *IntHeap)Pop() interface{} {
    x := (*h)[len(*h)-1]
    *h = (*h)[:(len(*h)-1)]
    return x
}

type KthLargest struct {
    kth *IntHeap
    k int
}


func Constructor(k int, nums []int) KthLargest {
    hp := &IntHeap{}
    heap.Init(hp)
    for _,n := range nums {
        heap.Push(hp,n)
        if hp.Len() > k {
            heap.Pop(hp)
        }
    }
    return KthLargest{hp,k}
}


func (this *KthLargest) Add(val int) int {
    heap.Push(this.kth,val)
    if this.kth.Len() > this.k {
        heap.Pop(this.kth)
    }
    return this.kth.Top()
}


/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
