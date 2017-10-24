/*
在英文字处理程序中，由于单词都是由字母序列构成，所以当输入到一行的末尾的时候，
就会遇到想要输入的单词长度大于所剩余的空白长度的情况，这就是折行问题。
对于手写文本，我们可以用连字符‘-’把单词分割到两行上，但是对于字处理程序而言，
其拥有更强的处理能力，可以通过运算来避免单词被分割到两行上

input:
        text: "I'm a good guy, and I know what I should not to do!"
        width: 25
output:
        how many lines. and text.


DP:http://blog.csdn.net/joylnwang/article/details/6775941
*/


package main

import (
        "fmt"
        "math"
)
const (
        kText1 = "I'm a good guy, and I know what I should not to do!"
)

func SplitWords(text string) []string {
        words := make([]string, 0)
        w := ""
        hasWord := false
        for _, t := range text {
                if t == ' ' {
                        if hasWord {
                                words = append(words, w)
                                hasWord = false
                                w = ""
                        }
                        continue
                }
                hasWord = true
                w = w + string(t)
        }
        if hasWord {
                words = append(words, w)
        }

        return words

}
func Greedy(text string, width int) (string, int) {
        words := SplitWords(text)
        context := ""
        idx := width
        lens := 1
        for _,w := range words {
                if idx - len(w) > 1 {
                        context += w + " "
                        idx -= len(w) + 1
                } else {
                        context += "\n" + w + " "
                        idx = width - len(w) - 1
                        lens += 1
                }
        }
        return context,lens
}


func Cost(words []string, width int, begin,end int) int {
        // 计算单行的cost,超行了是-1
        t := 0
        for i:=begin;i <= end;i++ {
                t += len(words[i])
        }
        cost := width - t - (end - begin)
        if cost > 0 {
                return cost *cost // 用方差作为代价,会让文本更收敛,并能保证行数最少
        }
        return -1
}
/*
        加入第j个单词后的最优结果
        f(j) = c(1,j) 没换行
        f(j) = min(f(k) + c(k,j)),换行了就把上一行的挪下来看看.
*/
func Dp(text string,width int) (string,int) {
        words := SplitWords(text)
        fmt.Println(words)
        // cost 一维数组, 记录的是f(j),从开头到j的代价
        cost := make([]int, len(words))
        // 当加入当前word后, 在本行的第一个元素的下标
        lines := make([]int, len(words))
        i,j,p,min := 0,0,0,0
        for ; i < len(words); i++ {
                t := Cost(words, width, 0, i)
                if t < 0 {
                        break
                } else {
                        cost[i] = t
                        lines[i] = 0
                }
        }
        fmt.Println(i)
        for ; i < len(words); i++ {
                min = int(math.MaxInt16)
                for j = i-1;j >0;j-- {
                        t := Cost(words,width,j+1,i)
                        if t >= 0 {
                                if t + cost[j] < min {
                                        min = t + cost[j]
                                        p = j
                                }
                        } else {
                                break
                        }
                }
                cost[i] = min
                lines[i] = p + 1
        }
        fmt.Println(lines)
        // 纠正
        i = len(words) - 1
        for i > 0 {
                j = lines[i]
                for i > j {
                        lines[i] = j
                        i = i-1
                }
                if i == 0 {
                        lines[i]=j
                } else {
                        lines[i] = j
                        i = i-1
                }
        }
        content := words[0] + " "
        num := 1
        for i = 1 ; i < len(words); i++ {
                if lines[i] != lines[i-1] {
                        content += "\n"
                        num += 1
                }
                content += words[i] + " "

        }

        fmt.Println(cost)
        fmt.Println(lines)
        return content,num
}

func main() {
        wds,len := Dp(kText1,25)
        fmt.Println(wds)
        fmt.Println(len)
        wds,len = Greedy(kText1,25)
        fmt.Println(wds)
        fmt.Println(len)
}
