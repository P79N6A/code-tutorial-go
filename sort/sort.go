package main

import (
        "sort"
        "fmt"
)

type Man struct {
        Age int
        Name string
}
type SortRRAttrs []Man

func (s SortRRAttrs) Len() int {
        return len(s)
}

func (s SortRRAttrs) Swap(i, j int) {
        s[i], s[j] = s[j], s[i]
}

func (s SortRRAttrs) Less(i, j int) bool {
        return (s[i].Age) >= (s[j].Age)
}

func main() {
        aa := []Man{Man{1,"1X"},Man{1,"2X"},Man{3,"3XX"},Man{3,"4XX"}}
        sort.Sort(SortRRAttrs(aa))
        fmt.Println(aa)
}
