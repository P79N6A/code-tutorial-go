package main

import "fmt"

// http://hustcat.github.io/array_and_slice/

func main() {
    arr := []int{1,2,3,4,5}
    arr = append(arr,6)
    copy(arr[1:],arr)
    fmt.Println(arr)
    
}
