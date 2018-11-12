package main

import "math"
/*
Given a non-negative integer c, your task is to decide whether there're two integers a and b such that a2 + b2 = c.

Example 1:
Input: 5
Output: True
Explanation: 1 * 1 + 2 * 2 = 5
Example 2:
Input: 3
Output: False
*/

func main() {
}
func judgeSquareSum(c int) bool {
    for i:=0;i*i<=c;i++ {
        if issquare(c-i*i) {
            return true
        }
    }
    return false
}
func issquare(num int) bool {
    //在一个有序的范围内找一个数=>应该能想到使用二分
    return binarysearch(0,num,num)
    // math.sqrt是ok的
    //x := int(math.Sqrt(float64(num)))
    //return x*x == num
    /*
    //下边的方法会超时.
    for i:=0;i*i<=num;i++ {
        if i*i == num {return true}
    }
    return false
    */
}
func binarysearch(start,end int, n int) bool {
    if start>end {return false}
    m := (start+end)/2
    if m*m == n {
        return true
    } else if m*m<n {
        return binarysearch(m + 1, end, n)
    }
    return binarysearch(start,m-1,n)
}