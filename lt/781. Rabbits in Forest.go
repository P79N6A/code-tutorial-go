package main

import "fmt"

/*
In a forest, each rabbit has some color. Some subset of rabbits (possibly all of them) tell you how many other rabbits have the same color as them. Those answers are placed in an array.

Return the minimum number of rabbits that could be in the forest.

Examples:
Input: answers = [1, 1, 2]
Output: 5
Explanation:
The two rabbits that answered "1" could both be the same color, say red.
The rabbit than answered "2" can't be red or the answers would be inconsistent.
Say the rabbit that answered "2" was blue.
Then there should be 2 other blue rabbits in the forest that didn't answer into the array.
The smallest possible number of rabbits in the forest is therefore 5: 3 that answered plus 2 that didn't.

Input: answers = [10, 10, 10]
Output: 11

Input: answers = []
Output: 0
Note:

answers will have length at most 1000.
Each answers[i] will be an integer in the range [0, 999].
*/

func main() {
    fmt.Println(numRabbits([]int{1,1,0,0,0}))
    fmt.Println(numRabbits([]int{2,2,2,2,2}))
}
func numRabbits(answers []int) int {
    ans := 0
    d := make(map[int]int)
    for _,other := range answers {
        if _,ok := d[other];!ok {
            // 第一次出现， other+1
            ans += 1 + other
            d[other]+=1
        } else if d[other] == other+1 {
            // 如果当前分组数量大于other的，说明兔子都出场了，清空了重来
            ans += 1+other
            d[other]=1
        } else {
            d[other]+=1
        }
    }
    return ans
}