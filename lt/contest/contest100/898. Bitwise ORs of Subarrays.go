package main

import "fmt"

func main() {
        fmt.Println(subarrayBitwiseORs([]int{1,1,2}))
        fmt.Println(subarrayBitwiseORs([]int{1,2,4}))
}

func subarrayBitwiseORs(A []int) int {
        retSet := make(map[int]bool) // all result
        tail := make(map[int]bool)  //  last result
        for i:=0;i<len(A);i++ {
                ntail := make(map[int]bool) // new tail result.
                for k,_ := range tail { //因为要遍历，所有需要new一个新的ntail
                        ntail[k|A[i]]=true
                        retSet[k|A[i]]=true
                }
                ntail[A[i]]=true // set itself.
                retSet[A[i]]=true // all result
                tail=ntail // swap tail, for next i
        }
        return len(retSet)
}