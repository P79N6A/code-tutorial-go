package main

import (
    "sort"
    "fmt"
)

func main() {
    pair := [][]int{{1,2},{2,3},{1,1},{4,5},{3,3}}

    sort.Slice(pair, func(i, j int) bool {
        return pair[i][0]<pair[j][0]
    })
    fmt.Println(pair)
    sort.Slice(pair, func(i, j int) bool {
        return pair[i][1]<pair[j][1]
    })
    fmt.Println(pair)

}
