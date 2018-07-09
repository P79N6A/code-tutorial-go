package main

import (
    "strings"
    "fmt"
)

/*

One way to serialize a binary tree is to use pre-order traversal. When we encounter a non-null node, we record the node's value. If it is a null node, we record using a sentinel value such as #.

     _9_
    /   \
   3     2
  / \   / \
 4   1  #  6
/ \ / \   / \
# # # #   # #
For example, the above binary tree can be serialized to the string "9,3,4,#,#,1,#,#,2,#,6,#,#", where # represents a null node.

Given a string of comma separated values, verify whether it is a correct preorder traversal serialization of a binary tree. Find an algorithm without reconstructing the tree.

Each comma separated value in the string must be either an integer or a character '#' representing null pointer.

You may assume that the input format is always valid, for example it could never contain two consecutive commas such as "1,,3".

Example 1:

Input: "9,3,4,#,#,1,#,#,2,#,6,#,#"
Output: true
Example 2:

Input: "1,#"
Output: false
Example 3:

Input: "9,#,#,1"
Output: false

*/
func isValidSerialization(preorder string) bool {
    idx := 0
    arr := strings.Split(preorder,",")
    if false == validtree(arr,&idx) {
        return false
    }
    if idx != len(arr)-1 {
        return false
    }
    return true
}
func validtree(data []string, idx *int) bool {
    if *idx >= len(data)  {
        return false
    }
    if data[*idx]=="#" {
        return true
    }
    *idx += 1
    left := validtree(data,idx)
    if left == false {
        return false
    }
    *idx += 1
    right := validtree(data,idx)
    if right == false {
        return false
    }
    return true
}
func main() {
    fmt.Println(isValidSerialization("9,#,#,1"))
    fmt.Println(isValidSerialization("9,#"))
    fmt.Println(isValidSerialization("9,3,4,#,#,1,#,#,2,#,6,#,#"))

}
