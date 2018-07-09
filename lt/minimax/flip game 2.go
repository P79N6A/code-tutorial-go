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
        for i:=1;i<len(str);i++ {
                if str[i-1]=='+' && str[i]=='+' {
                        substr := str[:i-1] + "--" + str[i+1:]
                        if !canWin(substr) {
                                return true
                        }
                }
        }
        return false
}


func main() {
        fmt.Println(canWin("++++"))
        fmt.Println(canWin("++"))
        fmt.Println(canWin("+-+"))

}
