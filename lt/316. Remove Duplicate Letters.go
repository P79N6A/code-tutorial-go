package main

import (
        "fmt"
        "sort"
)
/*

Given a string which contains only lowercase letters, remove duplicate letters so that every letter appear once and only once. You must make sure your result is the smallest in lexicographical order among all possible results.

Example 1:

Input: "bcabc"
Output: "abc"
Example 2:

Input: "cbacdcbc"
Output: "acdb"

*/
type lastEle struct {
        val byte
        idx int
}
type SortEle []lastEle
func (s SortEle)Swap(i,j int) {
        s[i],s[j]=s[j],s[i]
}
func (s SortEle)Less(i,j int) bool{
        return s[i].idx < s[j].idx
}
func (s SortEle)Len() int{
        return len(s)
}
func removeDuplicateLetters(s string) string {
        lastlet := make(map[byte]int)
        for i:=0;i<len(s);i++ {
                lastlet[s[i]] = i
        }
        lastele := make([]lastEle,0)
        for k,v := range lastlet {
                lastele = append(lastele,lastEle{k,v})
        }
        sort.Sort(SortEle(lastele))
        fmt.Println(lastele)
        lastp := 0
        ret := ""
        last := make(map[byte]bool)
        for _,l := range lastele {
                if last[l.val]==true {
                        lastp = l.idx
                        continue
                }
                c,pos := findmin(s,lastp+1,l.idx)
                if lastp == l.idx {
                        c = findmin(s,lastp,l.idx)
                }
                if c == l.val {
                        ret += string(c)
                } else {
                        last[c]=true
                }

                fmt.Println(string(c),lastp,ret,l)
                if last[c]==true {
                        lastp = l.idx
                        continue
                }
                fmt.Println(string(c),lastp,ret,l)
                ret += string(c)
                last[c]=true
                lastp = l.idx
        }
        return ret
}
func findmin(s string,i,j int) byte {
        var min  byte = 'z'
        for ii:=i;ii <= j&&ii < len(s);ii++ {
                if s[ii]<min {
                        min = s[ii]
                }
        }
        return min
}

func main() {
        fmt.Println(removeDuplicateLetters("cbacdcbc"))
        //fmt.Println(removeDuplicateLetters("bcabc"))
}
