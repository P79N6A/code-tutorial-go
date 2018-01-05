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

func doubleandcons(n int,list []int) []int{
        return append([]int{n * 2},list...)
}

func doubleall(list []int) []int{
        return foldr2(doubleandcons,make([]int,0),list)
}

func main() {

        inputList := []int{1, 2, 3}
        fmt.Printf("doubleall %v",doubleall(inputList))

}
