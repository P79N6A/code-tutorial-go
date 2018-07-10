package main

import "fmt"
/*
Given a word, you need to judge whether the usage of capitals in it is right or not.

We define the usage of capitals in a word to be right when one of the following cases holds:

All letters in this word are capitals, like "USA".
All letters in this word are not capitals, like "leetcode".
Only the first letter in this word is capital if it has more than one letter, like "Google".
Otherwise, we define that this word doesn't use capitals in a right way.
Example 1:
Input: "USA"
Output: True
Example 2:
Input: "FlaG"
Output: False
Note: The input will be a non-empty word consisting of uppercase and lowercase latin letters.


*/
func detectCapitalUse(word string) bool {
        numc := 0
        for i:=0;i<len(word);i++ {
                if word[i] <= 'Z' {
                        numc += 1
                }
        }
        return numc == 0 || (len(word)>1 && numc == 1 && word[0] <= 'Z') || numc == len(word)
}


func main() {
        //fmt.Println(detectCapitalUse("USA"))
        //fmt.Println(detectCapitalUse("A"))
        //fmt.Println(detectCapitalUse("Alice"))
        //fmt.Println(detectCapitalUse("alice"))
        fmt.Println(detectCapitalUse("FlaG"))

}
