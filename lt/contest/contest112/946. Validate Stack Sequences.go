package main
/*
Given two sequences pushed and popped with distinct values, return true if and only if this could have been the result of a sequence of push and pop operations on an initially empty stack.



Example 1:

Input: pushed = [1,2,3,4,5], popped = [4,5,3,2,1]
Output: true
Explanation: We might do the following sequence:
push(1), push(2), push(3), push(4), pop() -> 4,
push(5), pop() -> 5, pop() -> 3, pop() -> 2, pop() -> 1
Example 2:

Input: pushed = [1,2,3,4,5], popped = [4,3,5,1,2]
Output: false
Explanation: 1 cannot be popped before 2.


Note:

0 <= pushed.length == popped.length <= 1000
0 <= pushed[i], popped[i] < 1000
pushed is a permutation of popped.
pushed and popped have distinct values.
*/

import "fmt"

func main() {
    fmt.Println(validateStackSequences([]int{1,2,3,4,5},[]int{4,5,3,2,1}))
    fmt.Println(validateStackSequences([]int{1,2,3,4,5},[]int{4,3,5,1,2}))
    fmt.Println(validateStackSequences([]int{2,1,0},[]int{1,2,0}))
}

//根据push和pop 模拟下进出栈,看是否合理即可.
func validateStackSequences(pushed []int, poped []int) bool {
    stack := make([]int,0)
    pi := 0
    for i:=0;i<len(pushed);i++ {
        for len(stack) > 0 &&poped[pi]==stack[len(stack)-1] {
            // push的stack top 等于pop当前值,出栈
            pi += 1
            stack = stack[:(len(stack) - 1)]
        }
        stack = append(stack,pushed[i])
    }
    for len(stack)>0&&pi<len(poped) {
        if poped[pi]==stack[len(stack)-1] {
            pi += 1
            stack = stack[:(len(stack) - 1)]
        } else {
            return false
        }
    }
    return true
}