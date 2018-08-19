package main

import "fmt"

func main() {
        // mee aqq
        fmt.Println(findAndReplacePattern([]string{"abc","deq","mee","aqq","dkd","ccc"}, "abb"))
        //all
        fmt.Println(findAndReplacePattern([]string{"ef","fq","ao","at","lx"},"ya"))
}

func findAndReplacePattern(words []string, pattern string) []string {
        ret := make([]string,0)
        for _,word := range words {
                if len(word) != len(pattern) {continue}
                dict := make(map[byte]byte)
                match := true
                for i:=0;i<len(word);i++ {
                        /*
                        if _,ok := dict[word[i]];!ok {
                                dict[word[i]]=pattern[i]
                        } else {
                                if dict[word[i]] != pattern[i] {
                                        match = false
                                        continue
                                }
                        }
                        */
                        if _,ok := dict[pattern[i]];!ok {
                                dict[pattern[i]]=word[i]
                        } else {
                                if dict[pattern[i]] != word[i] {
                                        match = false
                                }
                        }
                }
                if match {
                        ret = append(ret,word)
                }
        }
        words = ret
        ret = make([]string,0)
        for _,word := range words {
                if len(word) != len(pattern) {continue}
                dict := make(map[byte]byte)
                match := true
                for i:=0;i<len(word);i++ {
                        if _,ok := dict[word[i]];!ok {
                                dict[word[i]]=pattern[i]
                        } else {
                                if dict[word[i]] != pattern[i] {
                                        match = false
                                        continue
                                }
                        }
                }
                if match {
                        ret = append(ret,word)
                }
        }
        return ret
}