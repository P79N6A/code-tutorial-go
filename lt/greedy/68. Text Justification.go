package main

import (
    "strings"
    "fmt"
)

/*
Given an array of words and a width maxWidth, format the text such that each line has exactly maxWidth characters and is fully (left and right) justified.

You should pack your words in a greedy approach; that is, pack as many words as you can in each line. Pad extra spaces ' ' when necessary so that each line has exactly maxWidth characters.

Extra spaces between words should be distributed as evenly as possible. If the number of spaces on a line do not divide evenly between words, the empty slots on the left will be assigned more spaces than the slots on the right.

For the last line of text, it should be left justified and no extra space is inserted between words.

Note:

A word is defined as a character sequence consisting of non-space characters only.
Each word's length is guaranteed to be greater than 0 and not exceed maxWidth.
The input array words contains at least one word.
Example 1:

Input:
words = ["This", "is", "an", "example", "of", "text", "justification."]
maxWidth = 16
Output:
[
   "This    is    an",
   "example  of text",
   "justification.  "
]
Example 2:

Input:
words = ["What","must","be","acknowledgment","shall","be"]
maxWidth = 16
Output:
[
  "What   must   be",
  "acknowledgment  ",
  "shall be        "
]
Explanation: Note that the last line is "shall be    " instead of "shall     be",
             because the last line must be left-justified instead of fully-justified.
             Note that the second line is also left-justified becase it contains only one word.
Example 3:

Input:
words = ["Science","is","what","we","understand","well","enough","to","explain",
         "to","a","computer.","Art","is","everything","else","we","do"]
maxWidth = 20
Output:
[
  "Science  is  what we",
  "understand      well",
  "enough to explain to",
  "a  computer.  Art is",
  "everything  else  we",
  "do                  "
]
*/

func main() {
    ret := fullJustify([]string{"ask","not","what","your","country","can","do","for","you","ask","what","you","can","do","for","your","country"},16)
    //ret := fullJustify([]string{"Science","is","what","we","understand","well","enough","to","explain", "to","a","computer.","Art","is","everything","else","we","do"},20)
    //ret := fullJustify([]string{"What","must","be","acknowledgment","shall","be"},16)
    for _,l := range ret {
        fmt.Println(l)
    }

}
func fullJustify(words []string, maxWidth int) []string {
    ret := make([]string,0)
    start,end := 0,0
    width := 0
    for end < len(words) {
        for end < len(words) {
            if width + len(words[end]) > maxWidth {
                break
            }
            width += len(words[end])+1 // one space
            end += 1
        }
        if end == len(words) {
            ret = append(ret,leftJustified(words[start:end],maxWidth))
        } else {
            ret = append(ret,fullyJustified(words[start:end],maxWidth))
        }
        start=end
        width = 0
    }
    return ret
}
func fullyJustified(words []string,width int) string {
    line := ""
    space := width
    for i:=0;i<len(words);i++ {
        space -= len(words[i])
    }
    if len(words) == 1 {
        line += words[0]
        line += strings.Repeat(" ",space)
    } else {
        for i := 0; i < len(words)-1; i++ {
            line += words[i]
            if i < space%(len(words)-1) {
                line += strings.Repeat(" ", space/(len(words)-1)+1)
            } else {
                line += strings.Repeat(" ", space/(len(words)-1))
            }
        }
        line += words[len(words)-1]
    }
    return line
}
func leftJustified(words []string,width int) string {
    space := width
    for _,w := range words {
        space -= len(w)
    }
    line := strings.Join(words," ")
    space -= len(words)-1
    if space > 0 {
        line += strings.Repeat(" ",space)
    }
    return line
}