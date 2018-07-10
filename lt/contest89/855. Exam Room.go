package main

import (
        "container/heap"
        "fmt"
)
/*
In an exam room, there are N seats in a single row, numbered 0, 1, 2, ..., N-1.

When a student enters the room, they must sit in the seat that maximizes the distance to the closest person.  If there are multiple such seats, they sit in the seat with the lowest number.  (Also, if no one is in the room, then the student sits at seat number 0.)

Return a class ExamRoom(int N) that exposes two functions: ExamRoom.seat() returning an int representing what seat the student sat in, and ExamRoom.leave(int p) representing that the student in seat number p now leaves the room.  It is guaranteed that any calls to ExamRoom.leave(p) have a student sitting in seat p.



Example 1:

Input: ["ExamRoom","seat","seat","seat","seat","leave","seat"], [[10],[],[],[],[],[4],[]]
Output: [null,0,9,4,2,null,5]
Explanation:
ExamRoom(10) -> null
seat() -> 0, no one is in the room, then the student sits at seat number 0.
seat() -> 9, the student sits at the last seat number 9.
seat() -> 4, the student sits at the last seat number 4.
seat() -> 2, the student​​​​​​​ sits at the last seat number 2.
leave(4) -> null
seat() -> 5, the student​​​​​​​ sits at the last seat number 5.
​​​​​​​

Note:

1 <= N <= 10^9
ExamRoom.seat() and ExamRoom.leave() will be called at most 10^4 times across all test cases.
Calls to ExamRoom.leave(p) are guaranteed to have a student currently sitting in seat number p.
*/
type pair struct {
        start int
        end int
}
type pairHeap []pair

func (h pairHeap) Len() int           { return len(h) }
func (h pairHeap) Less(i, j int) bool {
        if h[i].end+h[i].start == h[j].end+h[j].start {
                return h[i].start > h[j].start
        }
        return h[i].end+h[i].start < h[j].end+h[j].start
}
func (h pairHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *pairHeap) Push(x interface{}) {
        // Push and Pop use pointer receivers because they modify the slice's length,
        // not just its contents.
        *h = append(*h, x.(pair))
}

func (h *pairHeap) Pop() interface{} {
        old := *h
        n := len(old)
        x := old[n-1]
        *h = old[0 : n-1]
        return x
}


type ExamRoom struct {
        // key is start
        index map[int]int
        // key is end
        rindex map[int]int

        pairheap *pairHeap

        num int
        cap int
}


func Constructor(N int) ExamRoom {
       em := ExamRoom{
                index:make(map[int]int),
                rindex:make(map[int]int),
                pairheap:&pairHeap{pair{0,N-1}},
                cap:N,
        }
        heap.Init(em.pairheap)
        return em
}


func (this *ExamRoom) Seat() int {
        if this.num > this.cap {return -1}
        max := this.pairheap.Pop()
        maxp := max.(pair)
        pos := (maxp.end+maxp.start)/2
        this.index[pos]=maxp.end
        this.rindex[pos]=maxp.start
        this.pairheap.Push(pair{maxp.start,pos})
        this.pairheap.Push(pair{pos+1,maxp.end})
        return pos
}


func (this *ExamRoom) Leave(p int)  {
        if _,ok := this.rindex[p];!ok {
                return
        }
        start := this.rindex[p]
        end := this.index[p]
        this.pairheap.Push(pair{start,end})
        delete(this.index,p)
        delete(this.rindex,p)
}


func main() {
        em := Constructor(10)
        fmt.Println(em.Seat())
        fmt.Println(em.Seat())
        fmt.Println(em.Seat())
        fmt.Println(em.Seat())
}
