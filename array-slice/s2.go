package main

import "fmt"

func main() {
    aaa := []string{"a", "b", "c", "d", "e"}
    for i := 0; i < len(aaa); i++ {
        nss := make([]string, 0)
        nss = append(nss, aaa[i:]...)
        nss = append(nss, aaa[:i]...)
        fmt.Println(nss)
    }

    ai := []int{}
    ai = append(ai, 0)
    ai = append(ai, 0)
    ai = append(ai, 0)
    ai = append(ai, 0)
    ai = append(ai, 0)
    ai = append(ai, 0)
    fmt.Println(ai)
}
