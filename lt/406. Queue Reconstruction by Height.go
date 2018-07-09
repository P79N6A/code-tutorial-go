package main

import (
    "sort"
    "fmt"
)

/*

Suppose you have a random list of people standing in a queue. Each person is described by a pair of integers (h, k), where h is the height of the person and k is the number of people in front of this person who have a height greater than or equal to h. Write an algorithm to reconstruct the queue.

Note:
The number of people is less than 1,100.


Example

Input:
[[7,0], [4,4], [7,1], [5,0], [6,1], [5,2]]

Output:
[[5,0], [7,0], [5,2], [6,1], [4,4], [7,1]]

*/
type sortPeople [][]int
func (s sortPeople)Len() int {
    return len(s)
}
func (s sortPeople)Swap(i,j int) {
    s[i],s[j]=s[j],s[i]
}
func (s sortPeople)Less(i,j int)bool {
    if s[i][0]==s[j][0] {
        return s[i][1]<s[j][1]
    }
    return s[i][0]>s[j][0]

}

func reconstructQueue(people [][]int) [][]int {
    sort.Sort(sortPeople(people))
    fmt.Println(people)
    ret := make([][]int,0)
    for _,p := range people {
        ret = arrayAdd(ret,p[1],p)
    }
    return ret
}
func arrayAdd(ret [][]int,idx int,v[]int) [][]int {
    nret := make([][]int,0)
    if idx == 0 {
        nret = append(nret,v)
        nret = append(nret,ret...)
    } else if idx > len(ret)-1 {
        nret = append(nret,ret...)
        nret = append(nret,v)
    } else {
        nret = append(nret,ret[:idx]...)
        nret = append(nret,v)
        nret = append(nret,ret[idx:]...)
    }
    return nret
}
/*
身高h低的放在前边不影响k的排序
*/
func reconstructQueue(people [][]int) [][]int {
    // 第三个参数cap需要传入
    res := make([][]int, 0, len(people))

    // 按照 persons 的排序方式，对 people 进行排序
    sort.Sort(sortPeople(people))

    // 把 person 插入到 res[idx] 上
    insert := func(idx int, person []int) {
        // 不需要每次都make出来新的空间。。。 copy即可。
        res = append(res, person)
        // 插入到末尾
        if len(res)-1 == idx {
            return
        }
        // 插入到中间位置
        copy(res[idx+1:], res[idx:])
        res[idx] = person
    }

    for _, p := range people {
        insert(p[1], p)
    }

    return res
}


func main() {
    fmt.Println(reconstructQueue([][]int{
        []int{7,0},
        []int{4,4},
        []int{7,1},
        []int{5,0},
        []int{6,1},
        []int{5,2},
    }))
    
}
