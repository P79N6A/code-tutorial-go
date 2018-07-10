package main

import (
        "fmt"
)
/*
Strings A and B are K-similar (for some non-negative integer K) if we can swap the positions of two letters in A exactly K times so that the resulting string equals B.

Given two anagrams A and B, return the smallest K for which A and B are K-similar.

Example 1:

Input: A = "ab", B = "ba"
Output: 1
Example 2:

Input: A = "abc", B = "bca"
Output: 2
Example 3:

Input: A = "abac", B = "baca"
Output: 2
Example 4:

Input: A = "aabc", B = "abca"
Output: 2
Note:

1 <= A.length == B.length <= 20
A and B contain only lowercase letters from the set {'a', 'b', 'c', 'd', 'e'}


// solve2的递归是按照转换了来的,子问题针对string的大小没有变化.
调整到solve1,缩小string的大小看看.
*/
//////////////////////////////////////////////////////
func kSimilarity(A string, B string) int {
        return solve1([]byte(A), []byte(B), make(map[string]int))
}
func solve1(a []byte, b []byte, cache map[string]int) int {
        if string(a) == string(b) {
                return 0
        }
        key := string(a) + "-" + string(b)
        if cache[key] > 0 {
                return cache[key]
        }
        min := len(a)
        for i := 0; i < len(a); i++ {
                if a[i] == b[i] {
                        continue
                }
                for j := i + 1; j < len(a); j++ {
                        if a[i] != b[j] || a[j] == b[j] {
                                continue
                        }
                        a[i], a[j] = a[j], a[i]
                        s := solve1(a[i + 1:], b[i + 1:], cache)
                        if s < min {
                                min = s
                        }
                        a[i], a[j] = a[j], a[i]
                }
        }
        cache[key] = min + 1
        return min + 1
}
func kSimilarity2(A string, B string) int {
        return solve2([]byte(A), []byte(B), make(map[string]int), 0)
}
func solve2(a []byte, b []byte, cache map[string]int, i int) int {
        if string(a) == string(b) {
                return 0
        }
        key := string(a)
        if _, ok := cache[key]; ok {
                return cache[key]
        }
        min := len(a)
        for i < len(a) && a[i] == b[i] {
                i++
        }
        for j := i + 1; j < len(b); j++ {
                if a[j] == b[i] {
                        a[i], a[j] = a[j], a[i]
                        s := solve2(a, b, cache, i + 1)
                        if s + 1 < min {
                                min = s + 1
                        }
                        a[i], a[j] = a[j], a[i]
                }
        }
        cache[key] = min
        return min
}
func solve3(a []byte, b []byte, cache map[string]int, i int) int {
        if string(a) == string(b) {
                return 0
        }
        key := string(a)
        if _, ok := cache[key]; ok {
                return cache[key]
        }
        min := len(a)
        for i < len(a)  {
                if a[i] == b[i] {
                        i++
                        continue
                }
                for j := i + 1; j < len(b); j++ {
                        if a[j] == b[i] {
                                a[i], a[j] = a[j], a[i]
                                s := solve3(a, b, cache, i + 1)
                                if s + 1 < min {
                                        min = s + 1
                                }
                                a[i], a[j] = a[j], a[i]
                        }
                }
        }
        cache[key] = min
        return min
}

func main() {
        fmt.Println(kSimilarity2("aabc", "abca"))
        //fmt.Println(kSimilarity("abccaacceecdeea","bcaacceeccdeaae"))
        fmt.Println(kSimilarity2("cdebcdeadedaaaebfbcf", "baaddacfedebefdabecc"))
}
