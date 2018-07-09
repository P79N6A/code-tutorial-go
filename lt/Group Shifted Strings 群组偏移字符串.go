package main

import (
    "strconv"
    "fmt"
)

/*
Given a string, we can "shift" each of its letter to its successive letter, for example: "abc" -> "bcd". We can keep "shifting" which forms the sequence:

"abc" -> "bcd" -> ... -> "xyz"
Given a list of strings which contains only lowercase alphabets, group all strings that belong to the same shifting sequence.

For example, given: ["abc", "bcd", "acef", "xyz", "az", "ba", "a", "z"],
Return:

[
  ["abc","bcd","xyz"],
  ["az","ba"],
  ["acef"],
  ["a","z"]
]


Note: For the return value, each inner list's elements must follow the lexicographic order.


看差值是否相同 hashmap....
*/

func groupStrings(words []string) [][]string {
    h := make(map[string][]string)
    for _,w := range words {
        k := ""
        for i:=0;i<len(w);i++ {
            k += strconv.Itoa(int(w[0]-w[i]+26)%26) + ","
        }
        if len(h[k])<=0 {
            h[k] = make([]string,0)
        }
        h[k] = append(h[k],w)
    }
    fmt.Println(h)
    ret := make([][]string,0)
    for _,v := range h {
        ret = append(ret,v)
    }
    return ret
}

func main() {
    fmt.Println(groupStrings([]string{"abc", "bcd", "acef", "xyz", "az", "ba", "a", "z"}))
}
