package main

import (
    "sort"
    "fmt"
)

func main() {
    nums := []int{2,4,3,5,6,1,2,3}
    sort.Ints(nums[:5])
    fmt.Println(nums)
    
}
