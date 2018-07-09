package main

import "fmt"

/*
Given an array of characters, compress it in-place.

The length after compression must always be smaller than or equal to the original array.

Every element of the array should be a character (not int) of length 1.

After you are done modifying the input array in-place, return the new length of the array.


Follow up:
Could you solve it using only O(1) extra space?


Example 1:
Input:
["a","a","b","b","c","c","c"]

Output:
Return 6, and the first 6 characters of the input array should be: ["a","2","b","2","c","3"]

Explanation:
"aa" is replaced by "a2". "bb" is replaced by "b2". "ccc" is replaced by "c3".
Example 2:
Input:
["a"]

Output:
Return 1, and the first 1 characters of the input array should be: ["a"]

Explanation:
Nothing is replaced.
Example 3:
Input:
["a","b","b","b","b","b","b","b","b","b","b","b","b"]

Output:
Return 4, and the first 4 characters of the input array should be: ["a","b","1","2"].

Explanation:
Since the character "a" does not repeat, it is not compressed. "bbbbbbbbbbbb" is replaced by "b12".
Notice each digit has it's own entry in the array.
Note:
All characters have an ASCII value in [35, 126].
1 <= len(chars) <= 1000.
*/
func compress(chars []byte) int {
    p,q,r,clens := 0,0,0,len(chars)
    for ;q < clens; {
        for ;q < clens && chars[p] == chars[q]; {
            q += 1
        }
        num := q - p
        if num > 1 {
            chars[r] = chars[p]
            r += 1
            if num < 10 {
                chars[r] = '0' + byte(num)
                r += 1
            } else if num < 100 {
                chars[r] = '0' + byte(num/10)
                r += 1
                chars[r] = '0' + byte(num%10)
                r += 1
            } else if num < 1000 {
                chars[r] = '0' + byte(num/100)
                r += 1
                lnum := num % 100
                chars[r] = '0' + byte(lnum/10)
                r += 1
                chars[r] = '0' + byte(lnum%10)
                r += 1
            } else {
                chars[r] = '1'
                r += 1
                chars[r] = '0'
                r += 1
                chars[r] = '0'
                r += 1
                chars[r] = '0'
                r += 1
            }
        } else {
            chars[r]=chars[p]
            r += 1
        }
        p = q
    }
    num := q - p
    if num > 0 {
        if num > 1 {
            chars[r] = chars[p]
            r += 1
            if num < 10 {
                chars[r] = '0' + byte(num)
                r += 1
            } else if num < 100 {
                chars[r] = '0' + byte(num/10)
                r += 1
                chars[r] = '0' + byte(num%10)
                r += 1
            } else if num < 1000 {
                chars[r] = '0' + byte(num/100)
                r += 1
                lnum := num % 100
                chars[r] = '0' + byte(lnum/10)
                r += 1
                chars[r] = '0' + byte(lnum%10)
                r += 1
            } else {
                chars[r] = '1'
                r += 1
                chars[r] = '0'
                r += 1
                chars[r] = '0'
                r += 1
                chars[r] = '0'
                r += 1
            }
        } else {
            chars[r] = chars[p]
            r += 1
        }
    }
    return r
}

func main() {
    arr := []byte{'a','a','b','b','c','c','c'}
    fmt.Println(compress(arr))
    fmt.Println(string(arr))
}
