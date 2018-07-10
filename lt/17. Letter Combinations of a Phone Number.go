package main
/*
Given a string containing digits from 2-9 inclusive, return all possible letter combinations that the number could represent.

A mapping of digit to letters (just like on the telephone buttons) is given below. Note that 1 does not map to any letters.
Input: "23"
Output: ["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
*/
import "fmt"

func letterCombinations(digits string) []string {
        return showletter(digits)
}

var letters [8]string = [8]string{
        "abc", "def",
        "ghi", "jkl", "mno",
        "pqrs", "tuv", "wxyz",
}

func showletter(digits string) []string {
        if len(digits) <= 0 {
                return nil
        }
        ll := showletter(digits[1:])
        res := make([]string, 0)
        for _, l := range letters[int(digits[0] - '2')] {
                if ll != nil {
                        for _, lll := range ll {
                                res = append(res, string(l) + lll)
                        }
                } else {
                        res = append(res, string(l))
                }
        }
        return res
}

func main() {
        fmt.Println(letterCombinations("23"))
}
