package main

import "fmt"
const (
    k1 = 1<<iota
    k2
    k3
    k4
)

func f(grid [][]int) {
    grid[0][0]=3
}
func main() {
    fmt.Println(k1,k2,k3,k4)
    grid := [][]int{
        {0,1,2},
        {3,4,5},
    }
    f(grid)
    fmt.Println(grid)
    fmt.Println(10e5==1000000)
    fmt.Println(1e5)
}
