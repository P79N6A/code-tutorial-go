package main

import "fmt"
func main() {
	slice := []int{1,2,3,4,5}
	var sl2 = slice
	sl2[0] = 10
	fmt.Println(slice)
}
