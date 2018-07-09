package main

import (
        "fmt"
)

// mapper.

func foldr2(f func(a int,b []int) []int,z []int,list []int)[]int{
        if len(list) == 0{
                return z
        }
        return f(list[0],foldr2(f,z,list[1:]))
}

func double(n int,list []int) []int{
        return append([]int{n * 2},list...)
}

func filter(n int,list []int) []int{
        if n > 2 {
                return append([]int{n},list...)
        }
        return list
}


func main() {
        inputList := []int{1, 2, 3}
        fmt.Printf("double %v\n",foldr2(double,make([]int,0),inputList))
        fmt.Printf("filter %v\n",foldr2(filter,make([]int,0),inputList))


}
