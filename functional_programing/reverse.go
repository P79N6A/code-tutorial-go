package main

import "fmt"

func foldr3(f func(a,b []int) []int, z []int, list []int)[]int{
        if len(list) == 0{
                return z
        }
        return f(list[:1],foldr3(f,z,list[1:]))
}

func main() {
        ilist := []int{1,2,3,4,5,6,7,8}
        fn := func(a,b []int) []int {
                b = append(b,a...)
                return b
        }
        fmt.Printf("list %v\n",foldr3(fn,[]int{},ilist))
}
