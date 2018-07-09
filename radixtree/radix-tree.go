package main

import (
        "fmt"
        "github.com/hashicorp/go-immutable-radix"
        "runtime"
        "time"
)
func createTree() *iradix.Tree {
        // Create a tree
        r := iradix.New()
        r, _, _ = r.Insert([]byte(".fo.foo.com.xx"), 1)
        r, _, _ = r.Insert([]byte(".fo.foo.com"), 1)
        r, _, _ = r.Insert([]byte("bar."), 222)
        r, _, _ = r.Insert([]byte("fooz"), 2)
        r, _, _ = r.Insert([]byte("fooz"), 2)
        return r
}

func f2(r *iradix.Tree) {
        r.Root().Walk(func(k []byte, v interface{}) bool {
        //        fmt.Println(string(k))
        //        fmt.Println(v)
                return false
        })
        // Find the longest prefix match
        r.Root().LongestPrefix([]byte("foozip"))
        r.Root().LongestPrefix(nil)
        r.Root().LongestPrefix([]byte(""))
}
func fff() {
        current := runtime.GOMAXPROCS(runtime.NumCPU())
        fmt.Println(current)
        r := createTree()
        for i:=0;i<20;i++ {
                go func(r *iradix.Tree) {
                                f2(r)
                }(r)
        }
        for {
                time.Sleep(time.Second)
        }
}
func main() {
        r := createTree()
        a,b,c := r.Root().LongestPrefix([]byte("fox"))
        fmt.Println(a,b,c)
        r.Root().WalkPath([]byte(".fo.foo.com.xxx"), func(k []byte, v interface{}) bool {
                fmt.Println(string(k))
                return false
        })
        /*
        r.Root().WalkPrefix([]byte(".fo.foo"), func(k []byte, v interface{}) bool {

                fmt.Println("xxxxxxxxxxx")
                fmt.Println(string(k))
                return false
        })*/
}
