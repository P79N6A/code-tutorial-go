package main

import (
    "strconv"
    "fmt"
    "sort"
)

/*
Given a chemical formula (given as a string), return the count of each atom.

An atomic element always starts with an uppercase character, then zero or more lowercase letters, representing the name.

1 or more digits representing the count of that element may follow if the count is greater than 1. If the count is 1, no digits will follow. For example, H2O and H2O2 are possible, but H1O2 is impossible.

Two formulas concatenated together produce another formula. For example, H2O2He3Mg4 is also a formula.

A formula placed in parentheses, and a count (optionally added) is also a formula. For example, (H2O2) and (H2O2)3 are formulas.

Given a formula, output the count of all elements as a string in the following form: the first name (in sorted order), followed by its count (if that count is more than 1), followed by the second name (in sorted order), followed by its count (if that count is more than 1), and so on.

Example 1:
Input:
formula = "H2O"
Output: "H2O"
Explanation:
The count of elements are {'H': 2, 'O': 1}.
Example 2:
Input:
formula = "Mg(OH)2"
Output: "H2MgO2"
Explanation:
The count of elements are {'H': 2, 'Mg': 1, 'O': 2}.
Example 3:
Input:
formula = "K4(ON(SO3)2)2"
Output: "K4N2O14S4"
Explanation:
The count of elements are {'K': 4, 'N': 2, 'O': 14, 'S': 4}.
Note:

All atom names consist of lowercase letters, except for the first character which is uppercase.
The length of formula will be in the range [1, 1000].
formula will only consist of letters, digits, and round parentheses, and is a valid formula as defined in the problem.
*/
func countOfAtoms(formula string) string {
    dict := make(map[string]int)
    solve(1,formula,dict)
    keys := make([]string,0)
    for k,_ := range dict {
        keys = append(keys,k)
    }
    sort.Strings(keys)
    ret := ""
    for _,k := range keys {
        if dict[k] == 1 {
            ret += k
        } else {
            ret += fmt.Sprintf("%s%d",k,dict[k])
        }
    }

    return ret
}

func solve(factor int,str string,dict map[string]int) {
    pre,cur := 0,0
    for cur < len(str) {
        for cur<len(str)&&str[cur] != '(' {cur += 1}
        solveno(factor,str[pre:cur],dict)
        pre = cur
        l,r := 0,0
        for (l + r == 0 || l > r) && cur < len(str) {
            if str[cur]=='(' {l += 1}
            if str[cur]==')' {r += 1}
            cur += 1
        }
        if l + r > 0 {
            // get the number
            num := cur
            for ; num < len(str); num++ {
                if str[num] < '0' || str[num] > '9' {
                    break
                }
            }
            n, _ := strconv.Atoi(str[cur:num])
            solve(factor*n, str[(pre + 1):(cur - 1)], dict)
            cur = num
            pre = cur
        }
    }
}

func solveno(factor int,str string,dict map[string]int) {
    // Kg4N
    cur := 0
    atom := ""
    num := ""
    for cur < len(str) {
        if (str[cur] >= 'A' && str[cur] <= 'Z') || (str[cur] >= 'a' && str[cur] <= 'z' ){
            if atom != "" && (str[cur] >= 'A' && str[cur] <= 'Z') {
                if num == "" {
                    dict[atom] += factor
                } else {
                    n,_ := strconv.Atoi(num)
                    dict[atom] += n*factor
                }
                atom = ""
                num = ""
            }
            atom += string(str[cur])
        } else {
            // number
            num += string(str[cur])
        }
        cur += 1
    }
    if atom != "" {
        if num == "" {
            dict[atom] += factor
        } else {
            n,_ := strconv.Atoi(num)
            dict[atom] += n*factor
        }
    }
}

func main() {
    fmt.Println(countOfAtoms("Mg(OH)2"))
    fmt.Println(countOfAtoms("K4(ON(SO3)2)2"))
    fmt.Println(countOfAtoms("Ks"))
}
