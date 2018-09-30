package main

import (
    "strings"
    "fmt"
)

var doc=`

Given a C++ program, remove comments from it. The program source is an array where source[i] is the i-th line of the source code. This represents the result of splitting the original source code string by the newline character \n.

In C++, there are two types of comments, line comments, and block comments.

The string // denotes a line comment, which represents that it and rest of the characters to the right of it in the same line should be ignored.

The string /* denotes a block comment, which represents that all characters until the next (non-overlapping) occurrence of */ should be ignored. (Here, occurrences happen in reading order: line by line from left to right.) To be clear, the string /*/ does not yet end the block comment, as the ending would be overlapping the beginning.

The first effective comment takes precedence over others: if the string // occurs in a block comment, it is ignored. Similarly, if the string /* occurs in a line or block comment, it is also ignored.

If a certain line of code is empty after removing comments, you must not output that line: each string in the answer list will be non-empty.

There will be no control characters, single quote, or double quote characters. For example, source = "string s = "/* Not a comment. */";" will not be a test case. (Also, nothing else such as defines or macros will interfere with the comments.)

It is guaranteed that every open block comment will eventually be closed, so /* outside of a line or block comment always starts a new comment.

Finally, implicit newline characters can be deleted by block comments. Please see the examples below for details.

After removing the comments from the source code, return the source code in the same format.

Example 1:
Input:
source = ["/*Test program */", "int main()", "{ ", "  // variable declaration ", "int a, b, c;", "/* This is a test", "   multiline  ", "   comment for ", "   testing */", "a = b + c;", "}"]

The line by line code is visualized as below:
/*Test program */
int main()
{
// variable declaration
int a, b, c;
/* This is a test
   multiline
   comment for
   testing */
a = b + c;
}

Output: ["int main()","{ ","  ","int a, b, c;","a = b + c;","}"]

The line by line code is visualized as below:
int main()
{

int a, b, c;
a = b + c;
}

Explanation:
The string /* denotes a block comment, including line 1 and lines 6-9. The string // denotes line 4 as comments.
Example 2:
Input:
source = ["a/*comment", "line", "more_comment*/b"]
Output: ["ab"]
Explanation: The original source string is "a/*comment\nline\nmore_comment*/b", where we have bolded the newline characters.  After deletion, the implicit newline characters are deleted, leaving the string "ab", which when delimited by newline characters becomes ["ab"].
Note:

The length of source is in the range [1, 100].
The length of source[i] is in the range [0, 80].
Every open block comment is eventually closed.
There are no single-quote, double-quote, or control characters in the source code.

*/
`
func main() {
    //fmt.Println(removeComments([]string{"/*Test program //*/", "int main()", "{ ", "  // variable declaration ", "int a, b, c;", "/* This is a test", "   multiline  ", "   comment for ", "   testing */", "a = b + c;", "}"}))
    fmt.Println(removeComments([]string{"a/*comment", "line", "more_comment*/b"}))
    //fmt.Println(removeComments([]string{"struct Node{", "    /*/ declare members;/**/", "    int size;", "    /**/int val;", "};"}))
    fmt.Println(removeComments([]string{"main() {",
        "/* here is commments",
        "  // still comments */",
        "   double s = 33;",
        "   cout << s;",
        "}"}))
}
func removeComments(source []string) []string {
    comments := false
    lines := make([]string,0)
    ans := ""
    for _,line := range source {
        for i:=0;i<len(line);i++ {
            /*
            几种条件:在注释内部=>直到走出注释
                    注释外部，单行注释
                    开始多行注释过程
            */
            if comments {
                if line[i]=='*' && i < len(line)-1 && line[i+1]=='/' {
                    comments = false
                    i+=1
                }
            } else {
                if line[i]=='/' && i < len(line)-1 && line[i+1]=='*' {  // /*
                    comments = true
                    i+=1
                } else if line[i]=='/' && i < len(line)-1 && line[i+1]=='/' {  // //
                    break
                } else {
                    ans += string(line[i])
                }
            }
        }
        if !comments && len(ans) > 0 { //不在注释内部，并且ans不是空，说明一行有有效结果了，
            lines = append(lines,ans)
            ans = ""
        }
    }
    return lines
}
func removeComments1(source []string) []string {
    nosingle := make([]string,0)
    for _,s := range source {
        a := strings.Index(s,"/*")
        b := strings.Index(s,"//")
        if b < 0 {
            nosingle = append(nosingle,s)
        } else if a >= 0 {
            if b < a {
                nosingle = append(nosingle,s[:b])
            } else {
                nosingle = append(nosingle,s)
            }
        } else {
            nosingle = append(nosingle,s[:b])
        }
    }
    var magic byte = 'X'
    merge := make([]byte,0)
    for _,line := range nosingle {
        for i:=0;i<len(line);i++ {
            if line[i]==magic {
                merge = append(merge,line[i])
            }
            merge = append(merge,line[i])
        }
        merge = append(merge,magic)
    }
    merge = merge[:len(merge)-1]
    fmt.Println(string(merge))
    // split
    slim := []byte(merge)
    i,j := 0,0
    for {
        // /*/
        for ;i < len(slim)-1;i++ {
            if slim[i]=='/' && slim[i+1]=='*' {
                break
            }
        }
        if i >= len(slim)-1 {
            break
        }
        j = i+2
        for ;j<len(slim)-1;j++ {
            if slim[j]=='*' && slim[j+1]=='/' {
                copy(slim[i:],slim[(j+2):])
                slim = slim[:(len(slim)-(j+2-i))]
                break
            }
        }
    }

    fmt.Println(slim)
    splits := make([]string,0)
    split := make([]byte,0)
    for i:=0;i<len(slim);i++ {
        if slim[i] == magic {
            if i < len(slim)-1 && slim[i+1]!=magic {
                if len(split) > 0 {
                    splits = append(splits,string(split))
                }
                split = make([]byte,0)
            }
            //i+=1
        } else {
            split = append(split,slim[i])
        }
    }
    if len(split) > 0 {
        splits = append(splits,string(split))
    }
    return splits
}

