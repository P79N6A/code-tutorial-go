package main

import "fmt"

/*
Suppose you have N integers from 1 to N. We define a beautiful arrangement as an array that is constructed by these N numbers successfully if one of the following is true for the ith position (1 <= i <= N) in this array:

The number at the ith position is divisible by i.
i is divisible by the number at the ith position.
Now given N, how many beautiful arrangements can you construct?

Example 1:
Input: 2
Output: 2
Explanation:

The first beautiful arrangement is [1, 2]:

Number at the 1st position (i=1) is 1, and 1 is divisible by i (i=1).

Number at the 2nd position (i=2) is 2, and 2 is divisible by i (i=2).

The second beautiful arrangement is [2, 1]:

Number at the 1st position (i=1) is 2, and 2 is divisible by i (i=1).

Number at the 2nd position (i=2) is 1, and i (i=2) is divisible by 1.
Note:
N is a positive integer and will not exceed 15.
*/

func countArrangement(N int) int {
    num := make([]int,0)
    for i:=0;i<=N;i++ {num = append(num,i)}
    cnt := 0
    bk(num,1,&cnt)
    return cnt
}

func bk(data []int,start int, cnt *int) {
    if start == len(data) {
        *cnt += 1
        return
    }
    for i:=start;i<len(data);i++ {
        // 不允许当前的替换是无效的，因为替换后的位置可能再被替换。
        // 只保证当前是有效就可以。因为如果后续的无效，也不会进入到start==len(data)的结果集合中
        if (data[i]%start == 0 || start%data[i] == 0) {
            data[i],data[start]=data[start],data[i]
            bk(data,start+1,cnt)
            data[i],data[start]=data[start],data[i]
        }
    }
}


func main() {
    fmt.Println(countArrangement(6)) // 36
    //fmt.Println(countArrangement(2))
}
