package main

import "fmt"

func main() {
        fmt.Println(subarrayBitwiseORs([]int{1,1,2}))
        fmt.Println(subarrayBitwiseORs([]int{1,2,4}))
}

func subarrayBitwiseORs(A []int) int {
        retSet := make(map[int]int)
        tail := make(map[int]int)
        for i:=0;i<len(A);i++ {
                ntail := make(map[int]int)
                for k,_ := range tail {
                        ntail[k|A[i]]=1
                        retSet[k|A[i]]=1
                }
                ntail[A[i]]=1
                retSet[A[i]]=1
                tail=ntail
        }
        return len(retSet)
}