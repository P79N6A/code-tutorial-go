package main

import (
    "strings"
    "fmt"
)

/*
Given two strings A and B, find the minimum number of times A has to be repeated
such that B is a substring of it. If no such solution, return -1.

For example, with A = "abcd" and B = "cdabcdab".

Return 3, because by repeating A three times (“abcdabcdabcd”),
B is a substring of it; and B is not a substring of A repeated two times ("abcdabcd").

Note:
The length of A and B will be between 1 and 10000.
*/
func repeatedStringMatch(A string, B string) int {
    counter := 1
    ra := A
    for len(ra)<len(A)+len(A)+len(B) {
        if len(ra) < len(B) {
            ra += A
            counter += 1
        } else {
            if strings.Contains(ra, B) {
                return counter
            } else {
                ra += A
                counter += 1
            }
        }
        fmt.Println(ra)
    }
    return -1
}
func main() {
    fmt.Println(repeatedStringMatch("abcd","cdabcdab"))
    
}
