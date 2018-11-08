package main

import (
    "strconv"
    "fmt"
)

/*
You are playing the following Bulls and Cows game with your friend: You write down a number and ask your friend to guess what the number is. Each time your friend makes a guess, you provide a hint that indicates how many digits in said guess match your secret number exactly in both digit and position (called "bulls") and how many digits match the secret number but locate in the wrong position (called "cows"). Your friend will use successive guesses and hints to eventually derive the secret number.

Write a function to return a hint according to the secret number and friend's guess, use A to indicate the bulls and B to indicate the cows.

Please note that both secret number and friend's guess may contain duplicate digits.

Example 1:

Input: secret = "1807", guess = "7810"

Output: "1A3B"

Explanation: 1 bull and 3 cows. The bull is 8, the cows are 0, 1 and 7.
Example 2:

Input: secret = "1123", guess = "0111"

Output: "1A1B"

Explanation: The 1st 1 in friend's guess is a bull, the 2nd or 3rd 1 is a cow.
Note: You may assume that the secret number and your friend's guess only contain digits, and their lengths are always equal.
*/
func main() {
    fmt.Println(getHint("1123","0111"))
    fmt.Println(getHint("1807","7810"))
    fmt.Println(getHint("1122","2211")) //0A4B
}

/*
给定两个等长的字符串，字符串中只包含数字，求两个字符串中的bulls和cows。
其中，bulls表示在字符串同一位置的数值相同的数字的个数。
cows表示，guess中的数字在secret中出现了，但是处在不正确的位置上的数字的个数。
*/
/*
如果secret和guess中相同位置相同字符，bull += 1
剩下的记录下secret里面出现了多少字符。其实就是看secret里面能够抵消多少guess种的字符
在扫描一次，看位置相同，字符不同的时候，guess的字符能够呗secret的抵消掉
*/
func getHint(secret string, guess string) string {
    bull,cow := 0,0
    d := make(map[byte]int)
    for i:=0;i<len(secret);i++ {
        if secret[i]==guess[i] {
            bull += 1
        } else {
            d[secret[i]]+=1
        }
    }
    for i:=0;i<len(secret);i++ {
        if secret[i]!=guess[i] {
            if d[guess[i]]>0 {
                cow += 1
                d[guess[i]]-=1
            }
        }
    }
    return strconv.Itoa(bull)+"A"+strconv.Itoa(cow)+"B"
}