package main

import (
        "mustard/base/hash"
        "fmt"
)
func main() {
       // var sll = []byte{1,2,3,4,5}
        var sll = []byte{1,2}
        hash.ShuffleSub(&sll,0,len(sll))
        fmt.Println(sll)
}
