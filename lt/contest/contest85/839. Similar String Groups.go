package main

import (
        "fmt"
)
/*
Two strings X and Y are similar if we can swap two letters (in different positions) of X, so that it equals Y.

For example, "tars" and "rats" are similar (swapping at positions 0 and 2), and "rats" and "arts" are similar, but "star" is not similar to "tars", "rats", or "arts".

Together, these form two connected groups by similarity: {"tars", "rats", "arts"} and {"star"}.  Notice that "tars" and "arts" are in the same group even though they are not similar.  Formally, each group is such that a word is in the group if and only if it is similar to at least one other word in the group.

We are given a list A of unique strings.  Every string in A is an anagram of every other string in A.  How many groups are there?

Example 1:

Input: ["tars","rats","arts","star"]
Output: 2
Note:

A.length <= 2000
A[i].length <= 1000
A.length * A[i].length <= 20000
All words in A consist of lowercase letters only.
All words in A have the same length and are anagrams of each other.
The judging time limit has been increased for this question.

union find....
*/

func numSimilarGroups(A []string) int {
        // key is A[i],value is root
        groupSet := make(map[int]int)
        for i:=0;i <len(A);i++ {
                for j:=i+1;j < len(A);j++ {
                        if isconnect(A[i],A[j],i,j) {
                                tar := i
                                for {
                                        if _, ok := groupSet[tar]; ok && tar != groupSet[tar]{
                                                tar = groupSet[tar]
                                        } else {
                                                break
                                        }
                                }
                                src := j
                                for {
                                        if _, ok := groupSet[src]; ok && src != groupSet[src]{
                                                src = groupSet[src]
                                        } else {
                                                break
                                        }
                                }
                                if tar < src {
                                        groupSet[src] = tar
                                } else  {
                                        groupSet[tar] = src
                                }
                        }
                }
                fmt.Println(groupSet)
        }
        union := make(map[int]bool)
        for i:=0;i<len(A);i++ {
                p := i
                for {
                        if _,ok := groupSet[p];!ok || groupSet[p] == p {
                                break
                        }
                        p = groupSet[p]
                }
                union[p]=true
        }
        fmt.Println(union)
        return len(union)
}
func isconnect(str1,str2 string,i,j int) bool {
        if len(str1) != len(str2) {return false}
        if len(str1) <= 0 {return false}
        diff := make([]byte,0)
        for i:=0;i<len(str1);i++ {
                if str1[i] != str2[i] {
                        if len(diff) == 4 {
                                return false
                        }
                        diff = append(diff,str1[i])
                        diff = append(diff,str2[i])
                }
        }

        if diff[0] == diff[3] && diff[1] == diff[2] {
                fmt.Println(str1,str2,i,j)
                return true
        }
        return false
}

func main() {
        //fmt.Println(numSimilarGroups([]string{"tars","rats","arts","star"}))
        fmt.Println(numSimilarGroups([]string{"ufixvnxsdxalinayfaappbmmj","nxpxiaauvyjxasbfmfinmdpla","ujimiyxsaxpavnanfapmlbxdf","ufimvyxsaxpainanfapdlbxmj","nxpyajaumxixalbfvpdnmasfi","nxpxiaauvyjxpsbfmfinmdala","ufimvyxspxaainanfapdlbxmj","nxpyaiaumxjxapbfvlanmdsfi","ufimvyxspxapinanfaadlbxmj","nxpyaaauvxjxasbfmfinmdpli","nxpyajaumxixapbfvlanmdsfi","nxpyaaaumxjxasbfvfinmdpli","ufixvnbsdxalinayfamppxamj","ujimvyxsaxpainanfapdlbxmf","ufixvyxsdxalinanfaappbmmj","nxpyaaaumxjxapbfvlinmdsfi","ufixvyxspxalinanfaadpbmmj","nxpyaaaumxjxasbfvlinmdpfi","ufixvyxspxapinanfaadlbmmj","ufixvnbsdxalinayfaappxmmj","ufiavnbsdxxlinayfamppxamj","nxpyajaumxixapbfvldnmasfi","ufiavnbsdxxlinayfamppxajm","nxpyiaauvxjxasbfmfinmdpla","ujimiyxsaxpavnanfapdlbxmf"}))
        //fmt.Println(numSimilarGroups([]string{"qihcochwmglyiggvsqqfgjjxu","gcgqxiysqfqugmjgwclhjhovi","gjhoggxvcqlcsyifmqgqujwhi","wqoijxciuqlyghcvjhgsqfmgg","qshcoghwmglygqgviiqfjcjxu","jgcxqfqhuyimjglgihvcqsgow","qshcoghwmggylqgviiqfjcjxu","wcoijxqiuqlyghcvjhgsqgmgf","qshcoghwmglyiqgvigqfjcjxu","qgsjggjuiyihlqcxfovchqmwg","wcoijxjiuqlyghcvqhgsqgmgf","sijgumvhqwqioclcggxgyhfjq","lhogcgfqqihjuqsyicxgwmvgj","ijhoggxvcqlcsygfmqgqujwhi","qshcojhwmglyiqgvigqfgcjxu","wcoijxqiuqlyghcvjhgsqfmgg","qshcojhwmglyiggviqqfgcjxu","lhogcgqqfihjuqsyicxgwmvgj","xscjjyfiuglqigmgqwqghcvho","lhggcgfqqihjuqsyicxgwmvoj","lhgocgfqqihjuqsyicxgwmvgj","qihcojhwmglyiggvsqqfgcjxu","ojjycmqshgglwicfqguxvihgq","sijvumghqwqioclcggxgyhfjq","gglhhifwvqgqcoyumcgjjisqx"}))
       // fmt.Println(numSimilarGroups([]string{"ajdidocuyh","djdyaohuic","ddjyhuicoa","djdhaoyuic","ddjoiuycha","ddhoiuycja","ajdydocuih","ddjiouycha","ajdydohuic","ddjyouicha"}))
//        fmt.Println(numSimilarGroups([]string{"kccomwcgcs","socgcmcwkc","sgckwcmcoc","coswcmcgkc", "cowkccmsgc","cosgmccwkc","sgmkwcccoc","coswmccgkc", "kowcccmsgc","kgcomwcccs"}))

}
