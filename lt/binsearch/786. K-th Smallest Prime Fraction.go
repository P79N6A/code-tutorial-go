package main
/*
A sorted list A contains 1, plus some number of primes.  Then, for every p < q in the list, we consider the fraction p/q.

What is the K-th smallest fraction considered?  Return your answer as an array of ints, where answer[0] = p and answer[1] = q.

Examples:
Input: A = [1, 2, 3, 5], K = 3
Output: [2, 5]
Explanation:
The fractions to be considered in sorted order are:
1/5, 1/3, 2/5, 1/2, 3/5, 2/3.
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
    var (
        // 根据题意，应该是0-1之间的一个数字
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
func main() {
    
}
