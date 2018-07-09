package main

import (
    "strings"
    "math"
    "fmt"
)

/*

We are given N different types of stickers. Each sticker has a lowercase English word on it.

You would like to spell out the given target string by cutting individual letters from your collection of stickers and rearranging them.

You can use each sticker more than once if you want, and you have infinite quantities of each sticker.

What is the minimum number of stickers that you need to spell out the target? If the task is impossible, return -1.

Example 1:

Input:

["with", "example", "science"], "thehat"
Output:

3
Explanation:

We can use 2 "with" stickers, and 1 "example" sticker.
After cutting and rearrange the letters of those stickers, we can form the target "thehat".
Also, this is the minimum number of stickers necessary to form the target string.
Example 2:

Input:

["notice", "possible"], "basicbasic"
Output:

-1
Explanation:

We can't form the target "basicbasic" from cutting letters from the given stickers.
Note:

stickers has length in the range [1, 50].
stickers consists of lowercase English words (without apostrophes).
target has length in the range [1, 15], and consists of lowercase English letters.
In all test cases, all words were chosen randomly from the 1000 most common US English words, and the target was chosen as a concatenation of two random words.
The time limit may be more challenging than usual. It is expected that a 50 sticker test case can be solved within 35ms on average.
*/
func minStickers(stickers []string, target string) int {
    if hascontain(strings.Join(stickers,""),target)==false {
        return -1
    }
    min := math.MaxInt64
    bt(stickers,0,"",0,target,&min)
    return min
}
func bt(st []string,pos int,can string,num int,target string, min *int)  {
    fmt.Println(can)
    if containnum(can,target) {
        if num < *min {
            *min = num
        }
        return
    }
    for i:=pos;i<len(st);i++ {
        bt(st,i,can+st[i],num+1,target,min)
    }
}
func containnum(str,target string)bool {
    m := make(map[byte]int)
    for i:=0;i<len(str);i++ {
        m[str[i]]+=1
    }
    for i:=0;i<len(target);i++ {
        m[target[i]]-=1
        if m[target[i]] < 0 {
            return false
        }
    }
    return true
}
func hascontain(str,target string)bool {
    m := make(map[byte]bool)
    for i:=0;i<len(str);i++ {
        m[str[i]]=true
    }
    for i:=0;i<len(target);i++ {
        if m[target[i]]==false {
            return false
        }
    }
    return true
}
func main() {
    fmt.Println(minStickers([]string{"with", "example", "science"},"thehat"))
    
}
