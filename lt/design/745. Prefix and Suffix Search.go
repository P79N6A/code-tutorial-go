package main
/*

Given many words, words[i] has weight i.

Design a class WordFilter that supports one function, WordFilter.f(String prefix, String suffix). It will return the word with given prefix and suffix with maximum weight. If no word exists, return -1.

Examples:
Input:
WordFilter(["apple"])
WordFilter.f("a", "e") // returns 0
WordFilter.f("b", "") // returns -1
Note:
words has length in range [1, 15000].
For each test case, up to words.length queries WordFilter.f may be made.
words[i] has length in range [1, 10].
prefix, suffix have lengths in range [0, 10].
words[i] and prefix, suffix queries consist of lowercase letters only.

*/

func main() {
    
}
type WordFilter struct {
    // key: perfix + subfix,value:weight
    persubfix map[string]int
}


func Constructor(words []string) WordFilter {
    fix := make(map[string]int)
    for x:=0;x<len(words);x++ {
        word := words[x]
        // fetch all prefix suffix combine.
        for i:=0;i<=len(word);i++ {
            for j:=len(word);j>=0;j-- {
                key := word[:i] + "+" + word[j:]
                fix[key]=x
            }
        }
    }
    return WordFilter{fix}
}


func (this *WordFilter) F(prefix string, suffix string) int {
    if val,ok := this.persubfix[prefix+"+"+suffix];ok {
        return val
    }
    return -1
}


/**
 * Your WordFilter object will be instantiated and called as such:
 * obj := Constructor(words);
 * param_1 := obj.F(prefix,suffix);
 */