package main

import "sort"
import "fmt"

type ByLength []string

func (s ByLength) Len() int {
        return len(s)
}

func (s ByLength) Swap(i, j int) {
        s[i], s[j] = s[j], s[i]
}

func (s ByLength) Less(i, j int) bool {
        return len(s[i]) < len(s[j])
}
func sorttt(ss []string) {
        sort.Sort(ByLength(ss))
}
func main() {
        fruits := []string{"peach", "banana", "kiwi"}
        sorttt(fruits)
        fmt.Println(fruits)
        var aa []string = nil
        fmt.Println(aa)
        sort.Strings(aa)
        fmt.Println(aa)
}