package main

import (
    "strings"
    "sort"
    "errors"
    "fmt"
)

func main() {
    x := []string{"a1 9 2 3 1","g1 act car","zo4 4 7","ab1 off key dog","a8 act zoo"}
    x = reorderLogFiles(x)
    fmt.Println(x)
}
func isnum(str string) error {
    for i:=0;i<len(str);i++ {
        if str[i]<'0' || str[i]>'9' {
            return errors.New("letter")
        }
    }
    return nil
}
type logs []string
func (l logs)Len() int {
    return len(l)
}
func (l logs)Swap(i,j int) {
    l[i],l[j]=l[j],l[i]
}
func (l logs)Less(i,j int) bool {
    fsi := strings.Split(l[i]," ")
    erri := isnum(fsi[1])
    fsj := strings.Split(l[j]," ")
    errj := isnum(fsj[1])
    if erri!=nil {
        if errj!=nil {
            // all letter
            x := strings.Join(fsi[1:]," ")
            y := strings.Join(fsj[1:]," ")
            return x<y
        } else {
            // i letter j dig
            return true
        }
    } else {
        if errj!=nil {
            // i dig j letter
            return false
        } else {
            // all dig
            return i<j
        }
    }
}

func reorderLogFiles(ls []string) []string {
    nls := make([]string,0)
    lls := make([]string,0)
    for _,ll := range ls {
        fsi := strings.Split(ll," ")
        erri := isnum(fsi[1])
        if erri != nil {
            // letter
            lls = append(lls,ll)
        } else {
            nls = append(nls,ll)
        }
    }
    sort.Sort(logs(lls))
    lls = append(lls,nls...)
    return lls
}