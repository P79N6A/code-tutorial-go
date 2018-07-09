package main

import "fmt"

/*
Nearly every one have used the Multiplication Table.
But could you find out the k-th smallest number quickly from the multiplication table?

Given the height m and the length n of a m * n Multiplication Table,
and a positive integer k, you need to return the k-th smallest number in this table.

Example 1:
Input: m = 3, n = 3, k = 5
Output:
Explanation:
The Multiplication Table:
1	2	3
2	4	6
3	6	9

The 5-th smallest number is 3 (1, 2, 2, 3, 3).
Example 2:
Input: m = 2, n = 3, k = 6
Output:
Explanation:
The Multiplication Table:
1	2	3
2	4	6

The 6-th smallest number is 6 (1, 2, 2, 3, 4, 6).
Note:
The m and n will be in the range [1, 30000].
The k will be in the range [1, m * n]
*/
func findKthNumber(m int, n int, k int) int {
    // build matrix
    matrix := make([][]int,0)
    for i:=1;i<=m;i++ {
        m := make([]int,0)
        for j:=1;j<=n;j++ {
            m = append(m,i*j)
        }
        matrix = append(matrix,m)
    }
    fmt.Println(matrix)
    // binary search.
    l,r := matrix[0][0],matrix[m-1][n-1]
    for l < r {
        mid := l + (r-l)/2
        counter := 0
        for i:=1;i<=m;i++ {
            if mid / i > n {
                counter += n
            } else {
                counter += mid / i
            }
        }
        fmt.Println(mid,counter,l,r,k)
        if counter < k {
            l = mid + 1
        } else {
            r = mid
        }
    }
    return r
}

func findKthNumber2(m int, n int, k int) int {
    // binary search.
    l,r := 1,m*n
    for l < r {
        mid := l + (r-l)/2
        counter := 0
        for i:=1;i<=m;i++ {
            if mid / i > n {
                counter += n
            } else {
                counter += mid / i
            }
        }
        fmt.Println(mid,counter,l,r,k)
        if counter < k {
            l = mid + 1
        } else {
            r = mid
        }
    }
    return r
}

func main() {
    fmt.Println(findKthNumber(2,3,6))
}
