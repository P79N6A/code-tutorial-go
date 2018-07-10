package main

import "sort"
import "fmt"
/*
有需要随机查找,有需要修改位置. 好像array和linklist都不太完整....


TreeSet


*/

type ExamRoom struct {
        pos []int
        cap int
}


func Constructor(N int) ExamRoom {
        return ExamRoom{cap:N-1}
}


func (this *ExamRoom) Seat() int {
        if len(this.pos)<=0 {
                this.pos = append(this.pos,0)
                return 0
        }
        pos:=0
        max:=-1
        for i:=0;i<len(this.pos);i++ {
                if i == 0 && this.pos[0] != 0 {
                        if max < this.pos[0] {
                                max=this.pos[0]
                                pos = 0
                        }
                }
                if i == len(this.pos)-1 {
                        if this.cap - this.pos[i] > max {
                                pos = this.cap
                                max = this.cap-this.pos[i]
                        }
                } else {
                        l := (this.pos[i+1]-this.pos[i])/2
                        if l > max {
                                max = l
                                pos = (this.pos[i+1]+this.pos[i])/2
                        }
                }
        }
        this.pos = append(this.pos,pos)
        sort.Ints(this.pos)
        return pos
}


func (this *ExamRoom) Leave(p int)  {
        l,r := 0,len(this.pos)
        pos := -1
        for l < r {
                m := (l+r)/2
                if this.pos[m]== p {
                        pos = m
                        break
                } else if this.pos[m]<p {
                        l = m+1
                } else {
                        r=m
                }
        }
        if pos < 0 {
                return
        }
        if pos == 0{
                this.pos = this.pos[1:]
        } else if pos == len(this.pos)-1{
                this.pos = this.pos[:len(this.pos)-1]
        } else {
                npos := make([]int,0)
                npos = append(npos,this.pos[:pos]...)
                npos = append(npos,this.pos[pos+1:]...)
                this.pos=npos
        }
}


func main() {
        em := Constructor(8)
        fmt.Println(em.Seat())
        fmt.Println(em.Seat())
        fmt.Println(em.Seat())
        em.Leave(0)
        em.Leave(7)
        fmt.Print(em.pos)
        fmt.Println(em.Seat())
        fmt.Println(em.Seat())
        fmt.Println(em.Seat())
        fmt.Println(em.Seat())
        fmt.Println(em.Seat())
        fmt.Println(em.Seat())
}
