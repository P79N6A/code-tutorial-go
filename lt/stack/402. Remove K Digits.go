package main

import "fmt"

/*
Given a non-negative integer num represented as a string, remove k digits from the number so that the new number is the smallest possible.

Note:
The length of num is less than 10002 and will be ≥ k.
The given num does not contain any leading zero.
Example 1:

Input: num = "1432219", k = 3
Output: "1219"
Explanation: Remove the three digits 4, 3, and 2 to form the new number 1219 which is the smallest.
Example 2:

Input: num = "10200", k = 1
Output: "200"
Explanation: Remove the leading 1 and the number is 200. Note that the output must not contain leading zeroes.
Example 3:

Input: num = "10", k = 2
Output: "0"
Explanation: Remove all the digits from the number and it is left with nothing which is 0.

*/
func removeKdigitsByStack(num string, k int) string {
    /*
    题目的corner case还是挺多的
    1. 前导0，前导0的含义是，stack的长度大于1；如果最终只有一个结果并且是0的case可能会被包含
    2. 最后的结果是个0怎么办？ 比如  10 k=1 的case，最终输出应该是0
    2. k >= len(num)
    3. 重复的元素是否需要入栈？ 比如 112 k=1 的case
    4. 最终stack为空
    */
    stack := make([]byte,0)
    if len(num) <= k {return "0"}
    del := k
    for i:=0;i<len(num);i++ {
        for len(stack)>0 && num[i] < stack[len(stack)-1] && del > 0 {
            // 栈非空，并且当前数字比栈顶的小，并且还没删除完，说明该出栈了
            stack = stack[:len(stack)-1]
            del -= 1
        }
        if num[i] != '0' || len(stack) > 0 {
            // 防止前导0，这个也拒绝了只有一个0的case
            stack = append(stack, num[i])
        }
    }
    for del > 0 && len(stack) > 0 {
        // 有可能最终没有删干净，比如 123 k=1 最终stack里面就是123，没有人出栈，需要根据del把数字删干净
        stack = stack[:len(stack)-1]
        del -= 1
    }
    // 通过这句话保证了只有一个0的case可以通过。  10 k=1的case
    if len(stack) <= 0 {return "0"}
    return string(stack)
}
func removeKdigits(num string, k int) string {
    if len(num) <= k {return "0"}
    if k == 0 {
        return num
    }
    if num[1]=='0' {
        i:=1
        for ;i<len(num);i++ {
            if num[i]!='0'{break}
        }
        num := num[i:]
        return removeKdigits(num,k-1)
    }
    // first up.
    j:=1
    for ;j<len(num);j++ {
        if num[j]<num[j-1] {
            break
        }
    }
    j = j-1

    if j >= len(num) {
        num = num[:len(num)-1]
    } else if num[j]<num[0] {
        num = num[1:]
    } else {
        num = num[:j]+num[j+1:]
    }
    return removeKdigits(num,k-1)
}


func main() {
    fmt.Println(removeKdigits("1432219",4),removeKdigitsByStack("1432219",4))
    fmt.Println(removeKdigits("10200",1),removeKdigitsByStack("10200",1))
    fmt.Println(removeKdigits("10",2),removeKdigitsByStack("10",2))
    fmt.Println(removeKdigits("9",1),removeKdigitsByStack("9",1))
    fmt.Println(removeKdigits("5337",2),removeKdigitsByStack("5337",2))
    fmt.Println(removeKdigits("5332",3),removeKdigitsByStack("5332",3))
    fmt.Println(removeKdigits("1111111",4),removeKdigitsByStack("1111111",4))
    fmt.Println(removeKdigits("21",1),removeKdigitsByStack("21",1))
    fmt.Println(removeKdigits("121",1),removeKdigitsByStack("121",1))
    fmt.Println(removeKdigits("112",1),removeKdigitsByStack("112",1))
    fmt.Println(removeKdigits("178376",1),removeKdigitsByStack("178376",1))
    fmt.Println(removeKdigits("10",1),removeKdigitsByStack("10",1))
}
