package main
/*
Given an absolute path for a file (Unix-style), simplify it.

For example,
path = "/home/", => "/home"
path = "/a/./b/../../c/", => "/c"

Corner Cases:

Did you consider the case where path = "/../"?
In this case, you should return "/".
Another corner case is the path might contain multiple slashes '/' together, such as "/home//foo/".
In this case, you should ignore redundant slashes and return "/home/foo".
*/
import (
    "strings"
    "fmt"
)
func simplifyPath(path string) string {
    p := strings.Split(path,"/")
    ret := make([]string,0)
    for i:=0;i<len(p);i++ {
        if p[i]=="" || p[i]=="." {
            continue
        }
        if p[i] == ".." {
          if len(ret) > 0 {
              ret = ret[:len(ret)-1]
          }
        } else {
            ret=append(ret,p[i])
        }
    }
    return "/" + strings.Join(ret,"/")
}
func main() {
    fmt.Println(simplifyPath("/.."))
    fmt.Println(simplifyPath("/home/"))
    fmt.Println(simplifyPath("/a/./b/../../c/"))

}
