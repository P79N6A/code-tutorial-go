package main

import "fmt"

/*

Given n non-negative integers representing an elevation map where the width of each bar is 1, compute how much water it is able to trap after raining.


The above elevation map is represented by array [0,1,0,2,1,0,1,3,2,1,2,1]. In this case, 6 units of rain water (blue section) are being trapped. Thanks Marcos for contributing this image!

Example:

Input: [0,1,0,2,1,0,1,3,2,1,2,1]
Output: 6

有几种方法
1.竖着来
    遍历数组，从左到右一次，记录每个位置左边最大值；
    从右到左一次，记录每个位置右边最大值
    在遍历数组高度，如果当前高度小于左右的最大值，说明有坑，高度就是左右最大值较小的。将这个位置的一竖条算到结果中去

2.横着来
    计算最大高度，按照高度遍历
    先找到第一个大于高度的左边。然后再进来比高度大的，说明可以有坑，减去左右放到结果，在继续过程
    重复其他高度

*/
func trap(height []int) int {
    lh := make([]int,len(height))
    rh := make([]int,len(height))
    max := 0
    for i:=0;i<len(height);i++ {
        if max < height[i] {max = height[i]}
        lh[i]=max
    }
    max = 0
    for i:=len(height)-1;i>=0;i-- {
        if max < height[i] {max = height[i]}
        rh[i]=max
    }
    t := 0
    for i:=0;i<len(height);i++ {
        min := lh[i]
        if rh[i] < min {
            min = rh[i]
        }
        if min > height[i] {
            t += min-height[i]
        }
    }
    return t
}
func main() {
    fmt.Println(trap([]int{0,1,0,2,1,0,1,3,2,1,2,1}))
    
}
