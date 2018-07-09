package main

import "fmt"

func bt(data []int, pos int, res *[]int, ret *[][]int) {
        if len(*res) == 2 {
                n := make([]int, len((*res)))
                copy(n, *res)
                *ret = append(*ret, n)
                return
        }
        if pos >= len(data) {
                return
        }
        bt(data, pos + 1, res, ret)
        *res = append(*res, data[pos])
        bt(data, pos + 1, res, ret)
        *res = (*res)[:len(*res) - 1]
}

func main() {
        data := []int{8, 4, 2, 1}
        ret := make([][]int, 0)
        res := make([]int, 0)
        bt(data, 0, &res, &ret)
        fmt.Println(ret)
}
