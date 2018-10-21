package main
/*
Given an array A of 0s and 1s, divide the array into 3 non-empty parts such that all of these parts represent the same binary value.

If it is possible, return any [i, j] with i+1 < j, such that:

A[0], A[1], ..., A[i] is the first part;
A[i+1], A[i+2], ..., A[j-1] is the second part, and
A[j], A[j+1], ..., A[A.length - 1] is the third part.
All three parts have equal binary value.
If it is not possible, return [-1, -1].

Note that the entire part is used when considering what binary value it represents.  For example, [1,1,0] represents 6 in decimal, not 3.  Also, leading zeros are allowed, so [0,1,1] and [1,1] represent the same value.



Example 1:

Input: [1,0,1,0,1]
Output: [0,3]
Example 2:

Input: [1,1,0,1,1]
Output: [-1,-1]


Note:

3 <= A.length <= 30000
A[i] == 0 or A[i] == 1

*/

import (
        "fmt"
)

func main() {
        /*
        fmt.Println(threeEqualParts([]int{0,0,0,1,0,1,0,1,0,1,1,0,1}))
        fmt.Println(threeEqualParts([]int{1,0,1,0,1}))//0,3
        fmt.Println(threeEqualParts([]int{1,1,0,1,1}))//-1,-1
        fmt.Println(threeEqualParts([]int{0,0,0,0,0}))//0,4
        fmt.Println(threeEqualParts([]int{1,1,0,0,1})) //0,2
        */
        fmt.Println(threeEqualParts([]int{0,1,1,0,0,1,1,1,1,1})) //-1,-1
}
/*
边界情况较多. 需要思路清晰,前导0是无效的,后面0是ok的.
解法:
1.先确定第三个区间.因为后缀0有效,找到第三个区间第一个是1的,代码中的third
2.找第一个区间.逻辑是,去掉前导0后,跟third相同即可.
3.从1,2获得了一个有效第二区间,从nozero+len(third) 到 len(A)-len(third)
这部分也去掉前导0,看剩余的是否有相同的前缀,并且不同的部分,全都是0
*/
func threeEqualParts(A []int) []int {
        nozero := 0
        for i:=0;i<len(A)&&A[i]==0;i++ {
                nozero += 1
        }
        if nozero == len(A) {return []int{0,len(A)-1}} // no 1
        cnt1 := 0
        for i:=0;i<len(A);i++ {
                if A[i]==1{cnt1+=1}
        }
        //if cnt1%3!=0 {return []int{-1,-1}}

        for i:=len(A)-1;i>(len(A)-1-len(A)/3);i-- {
                if A[i] == 0 {continue} // only consider startwith 1
                third := A[i:]
                if !hasSubfixOnlyZero(nozero,A,third) {continue}
                //fmt.Println(third,A[(nozero+len(third)):i])
                idx := containOnlyZero(nozero+len(third),i,A,third)
                if idx >= 0 {
                        return []int{nozero+len(third)-1,idx}
                }
        }
        return []int{-1,-1}
}
func containOnlyZero(start,end int,A []int,third []int) int {
        // 去掉A的前导0,看A=third + 0
        i := start
        for i=start;i<len(A)&&A[i]==0;i++ {
                start += 1
        }
        for i=start;i<start+len(third);i++ {
                if A[i]!=third[i-start] {
                        return -1
                }
        }
        idx := i
        for ;i<end;i++ {
                if A[i]==1{return -1}
        }
        return idx
}
func hasSubfixOnlyZero(nozero int, A []int, third []int) bool {
        for i:=nozero;i<len(third)+nozero;i++ {
                if A[i]!=third[i-nozero] {
                        return false
                }
        }
        return true
}
////////////////
func threeEqualParts1(A []int) []int {

        pre0 := 0
        for i:=0;i<len(A)&&A[i]==0;i++ {
                pre0 += 1
        }
        A = A[pre0:]
        for i:=1;i<=(pre0+len(A))/3&&i<len(A);i++ {
                idx,same := hasSubfix(A,i)
                if idx < 0 {
                        continue
                }
                //fmt.Println(A[:i],same,A[i:idx])
                leftz := hasprefixOnlyZero(same,A[i:idx])
                if leftz >= 0 {
                        return []int{pre0+i-1,pre0+idx-leftz}
                }
        }
        return []int{-1,-1}
}
func hasprefixOnlyZero(A []int, B []int) int {
        if len(A) > len(B) {return -1}
        a,b := 0,0
        for i:=0;i<len(B)&&B[i]==0;i++ {b += 1}
        for ;a<len(A)&&b<len(B)&&A[a]==B[b];a,b=a+1,b+1 {}
        left:=len(B)-b
        for ;b<len(B);b++ {
                if B[b]==1 {return -1}
        }
        return left

}
func hasSubfix(A []int, ii int) (int,[]int) {
        for i:=0;i<ii;i++ {
                if A[i]!=A[len(A)-ii+i] {
                        return -1,nil
                }
        }
        return len(A)-ii,A[:ii]
}