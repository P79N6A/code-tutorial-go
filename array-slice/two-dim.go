package main

import "fmt"

func main() {
        array := make([][]int,10)
        for i := 0;i < 10;i++ {
                array[i] = make([]int,5)
                for j := 0;j < 5;j++{
                        array[i][j]=i * j
                }
        }
        fmt.Println(array)

        arr := [2][3]int{{1, 2, 3}, {4, 5, 6}}
        for i := range arr {
                for j := range arr[i] {
                        fmt.Printf("%v ", arr[i][j])
                }
                fmt.Println()
        }
        arr2 := [2][3]int{}
        fmt.Println(arr2)
        for i:=0;i<2;i++ {
                for j:=0;j<3;j++ {
                        fmt.Println(i,j,arr2[i][j])
                }
        }
}
