package main

import "fmt"

/*
In the "100 game," two players take turns adding, to a running total, any integer from 1..10. The player who first causes the running total to reach or exceed 100 wins.

What if we change the game so that players cannot re-use integers?

For example, two players might take turns drawing from a common pool of numbers of 1..15 without replacement until they reach a total >= 100.

Given an integer maxChoosableInteger and another integer desiredTotal, determine if the first player to move can force a win, assuming both players play optimally.

You can always assume that maxChoosableInteger will not be larger than 20 and desiredTotal will not be larger than 300.

Example

Input:
maxChoosableInteger = 10
desiredTotal = 11

Output:
false

Explanation:
No matter which integer the first player choose, the first player will lose.
The first player can choose an integer from 1 up to 10.
If the first player choose 1, the second player can only choose integers from 2 up to 10.
The second player will win by choosing 10 and get a total = 11, which is >= desiredTotal.
Same with other integers chosen by the first player, the second player will always win.
*/
func canIWin(maxChoosableInteger int, desiredTotal int) bool {
    // 1...maxChoosableInteger
    // corner case: 如果第一次就能达到目标，肯定是true
    if maxChoosableInteger >= desiredTotal {return true}
    // corner case: 如果把所有的数字都拿了还是不能达到目标，就是false
    if maxChoosableInteger*(maxChoosableInteger+1)/2 < desiredTotal {return false}
    return canWin(maxChoosableInteger,desiredTotal,0,make(map[int]bool))
}
func canWin(maxChoosableInterger,desiredTotal int, used int,cache map[int]bool) bool {
    if desiredTotal <= 0 {
        // 如果目标小于等于0了，说明上次是对方赢了
        // 所有的出口位置都需要添加缓存
        /*
        因为用的是bitmap，这样cache的key就是bitmap，相对hashmap有较大优势。bitmap的存储功能
        */
        cache[used]=false
        return false
    }
    if _,ok := cache[used];ok {
        return cache[used]
    }
    for i:=1;i<=maxChoosableInterger;i++ {
        cur := 1 << uint(i)
        if cur & used == 0 {
            // bool值的MiniMax算法，递归过程需要交替的取反进行
            // 因为true是只要能赢就可以
            // false是所有的情况都不能赢，更严格
            if !canWin(maxChoosableInterger,desiredTotal-i,used|cur,cache) {
                cache[used]=true
                return true
            }
        }
    }
    // 所有的都不能赢才是输了
    cache[used]=false
    return false
}
func main() {
    fmt.Println(canIWin(4,6))
    
}
