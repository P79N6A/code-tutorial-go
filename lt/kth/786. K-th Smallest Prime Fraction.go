package main

import "fmt"

/*
A sorted list A contains 1, plus some number of primes.
Then, for every p < q in the list, we consider the fraction p/q.

What is the K-th smallest fraction considered?
Return your answer as an array of ints, where answer[0] = p and answer[1] = q.

Examples:
Input: A = [1, 2, 3, 5], K = 3
Output: [2, 5]
Explanation:
The fractions to be considered in sorted order are:
1/5, 1/3, 2/5, 1/2, 3/5, 2/3.

1/2 1/3 1/3
2/3 2/5
3/5

The third fraction is 2/5.

Input: A = [1, 7], K = 1
Output: [1, 7]
Note:

A will have length between 2 and 2000.
Each A[i] will be between 1 and 30000.
K will be between 1 and A.length * (A.length - 1) / 2.

*/

func kthSmallestPrimeFraction(A []int, K int) []int {
    alen := len(A)
    if alen <= 2 {return A}
    var (
        l float64 = 0
        r float64 = 1
    )
    pl,pr := 0,1
    for l < r {
        mid := (l + r)/2
        counter,pl := 0,0
        for i,j:=0,alen-1;i<alen;i++ {
            for ;j>=0&&float64(A[i]) > mid * float64(A[alen-1-j]);j-- {}
            counter += j+1
            if j >=0&&pl*A[alen-1-j] < pr*A[i] {
                // 多行处理记录下最小最接近的数
                // pl/pr < A[i]/A[j]
                pl,pr = A[i],A[alen-1-j]
            }
        }
        if counter < K {
            l = mid
        } else if counter > K{
            r = mid
        } else {
            return []int{pl,pr}
        }
    }
    return []int{pl,pr}
}

func kthSmallestPrimeFraction2(A []int, K int) []int {
    /*
    1/2  1/3  1/5
         2/3  2/5
              3/5
    从左到右是降序的
    问题是找第k小
    找一个p,q记录查找过程中遇到的比a[i]/a[j]小的最大值； 最接近他的。
    p/q < A[i]/A[j]
    */
    alen := len(A)
    var (
        l float64 = 0
        r float64 = 1
    )
    p,q := 0,1
    for l < r {
        mid := (l + r)/2
        counter:= 0
        for i:=0;i<alen-1;i++ {
            for j:=alen-1;j>i;j-- {
                if float64(A[i]) < mid * float64(A[j]) {
                    counter++
                    if p*A[j] < q*A[i] { // 所有有效范围内的最大值。
                        p, q = A[i], A[j]
                    }
                }
            }
            fmt.Println(l,r,mid,counter,p,q)
        }
        if counter < K {
            l = mid
        } else if counter > K{
            r = mid
        } else {
            return []int{p,q}
        }
        p,q = 0,1
    }
    return []int{p,q}
}

func main() {
    fmt.Println(kthSmallestPrimeFraction2([]int{1,2,3,5},3))
}
