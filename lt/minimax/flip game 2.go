package main

import "fmt"
/*
[LeetCode] Flip Game 翻转游戏之二


You are playing the following Flip Game with your friend: Given a string that contains only these two characters: + and -, you and your friend take turns to flip twoconsecutive "++" into "--". The game ends when a person can no longer make a move and therefore the other person will be the winner.

Write a function to determine if the starting player can guarantee a win.

For example, given s = "++++", return true. The starting player can guarantee a win by flipping the middle "++" to become "+--+".

Follow up:
Derive your algorithm's runtime complexity.


*/

func canWin(str string) bool {
        /*
        只要有一种方法能赢就是可以赢[返回true的时机]
        只有所有的方法都不能赢才是输了[返回false的时机]
        */
        for i:=1;i<len(str);i++ {
                if str[i-1]=='+' && str[i]=='+' {
                        substr := str[:i-1] + "--" + str[i+1:]
                        // 取反，变成了严格的返回
                        if !canWin(substr) {
                                return true
                        }
                }
        }
        // 对于不能赢是严格的，所有的尝试都不能成功才是失败。
        // 因为对手都是聪明的，会使用最优的结果
        return false
}


func main() {
        fmt.Println(canWin("++++"))
        fmt.Println(canWin("++"))
        fmt.Println(canWin("+-+"))

}
