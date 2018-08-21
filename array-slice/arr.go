package main

import "fmt"

func f(grid [][]int) {
    grid[0][0]=3
}
func main() {
    grid := [][]int{
        {0,1,2},
        {3,4,5},
    }
    f(grid)
    fmt.Println(grid)
}
