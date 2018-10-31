package main

import "fmt"

func main() {
	fmt.Println(1 ^ 1)
	fmt.Println(0 ^ 1)
	fmt.Println(1 ^ 0)
	fmt.Println(0 ^ 0)
	x := 9
	fmt.Println(^x + 1)
	fmt.Println(^(x - 1))
	fmt.Println(x & (-x))

}
