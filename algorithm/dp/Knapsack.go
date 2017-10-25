/*
数组中的数分为两组，给出一个算法，使得两个组的和的差的绝对值最小数组中的数的取值范围是0<x<100，元素个数也是大于0，小于100
比如a[]={2,4,5,6,7},得出的两组数{2，4,，6}和{5，7}，abs(sum(a1)-sum(a2))=0;
比如{2，5，6，10}，abs(sum(2,10)-sum(5,6))=1,所以得出的两组数分别为{2，10}和{5,，6}。

http://karaffeltut.com/NEWKaraffeltutCom/Knapsack/knapsack.html

TODO: get path... 非递归
主要就是逆序判断元素是否被选中
*/
package main

import "fmt"

/*
如果k放进去没有剩余容量
        f(i,w) = f(i-1,w)
如果k放进去有剩余容量
        f(i,w) = f(i-1,w-wi) + vi // pick
        f(i,w) = f(i-1,w)  // no pick
*/
func DP(array []int,k,left int,result *[]int) int {
        // input : k, 第k个物品
        // input : left 剩余容量
        // output : 最大价值
        if k < 0 {
                return 0
        }
        if array[k] > left {
                (*result)[k] = 0
                return DP(array,k-1,left,result)
        } else {
                pick := DP(array,k-1,left - array[k],result) + array[k]
                nopick := DP(array,k-1,left,result)
                if pick > nopick {
                        (*result)[k] = 1
                        return pick
                } else {
                        (*result)[k] = 0
                        return nopick
                }
        }
}

func main() {
        result := make([]int,7)
        ret := DP([]int{9,3,4,10,5,6,4},6,20,&result)
        fmt.Println(ret)
        fmt.Println(result)
}
