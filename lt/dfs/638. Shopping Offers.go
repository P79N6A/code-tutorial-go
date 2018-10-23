package main

import "fmt"

/*
In LeetCode Store, there are some kinds of items to sell. Each item has a price.

However, there are some special offers, and a special offer consists of one or more different kinds of items with a sale price.

You are given the each item's price, a set of special offers, and the number we need to buy for each item. The job is to output the lowest price you have to pay for exactly certain items as given, where you could make optimal use of the special offers.

Each special offer is represented in the form of an array, the last number represents the price you need to pay for this special offer, other numbers represents how many specific items you could get if you buy this offer.

You could use any of special offers as many times as you want.

Example 1:
Input: [2,5], [[3,0,5],[1,2,10]], [3,2]
Output: 14
Explanation:
There are two kinds of items, A and B. Their prices are $2 and $5 respectively.
In special offer 1, you can pay $5 for 3A and 0B
In special offer 2, you can pay $10 for 1A and 2B.
You need to buy 3A and 2B, so you may pay $10 for 1A and 2B (special offer #2), and $4 for 2A.
Example 2:
Input: [2,3,4], [[1,1,0,4],[2,2,1,9]], [1,2,1]
Output: 11
Explanation:
The price of A is $2, and $3 for B, $4 for C.
You may pay $4 for 1A and 1B, and $9 for 2A ,2B and 1C.
You need to buy 1A ,2B and 1C, so you may pay $4 for 1A and 1B (special offer #1), and $3 for 1B, $4 for 1C.
You cannot add more items, though only $9 for 2A ,2B and 1C.
Note:
There are at most 6 kinds of items, 100 special offers.
For each item, you need to buy at most 6 of them.
You are not allowed to buy more items than you want, even if that would lower the overall price.
*/

func main() {
    fmt.Println(shoppingOffers([]int{2,3,4}, [][]int{{1,1,0,4},{2,2,1,9}}, []int{1,2,1}))
    fmt.Println(shoppingOffers([]int{2,5}, [][]int{{3,0,5},{1,2,10}}, []int{3,2}))

}

func shoppingOffers(price []int, special [][]int, needs []int) int {
    return dfs(needs,special,price)
}
func dfs(u []int,special [][]int,price []int) int {
    // return min cost.
    min := directCost(u,price) // 直接购买，不组合的花费
    for _,s := range special {
        ok,v := next(u,s)
        if !ok {continue}
        // 使用各个有效的组合方式后，能否将价格降低。
        x := dfs(v,special,price)+s[len(u)]
        if x < min {min=x}
    }
    return min
}
func next(u []int,s []int) (bool,[]int) {
    r := make([]int,len(u))
    for i:=0;i<len(u);i++ {
        r[i]=u[i]-s[i]
        if r[i]<0{return false,nil}
    }
    return true,r
}
func directCost(u []int, price []int) int {
    ans := 0
    for i:=0;i<len(u);i++ {
        ans += u[i]*price[i]
    }
    return ans
}